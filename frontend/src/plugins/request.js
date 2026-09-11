import axios from "axios";
import { apiRoutes } from "./apiRoutes.js";
import { innerRoutes } from "./routes.js";

const api = axios.create({
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
});

let isRefreshing = false;
let failedQueue = [];
let isRefreshingFailed = false;

const processQueue = (error, token = null) => {
    failedQueue.forEach(prom => {
        if (error) {
            prom.reject(error);
        } else {
            prom.resolve(token);
        }
    });
    failedQueue = [];
};

api.interceptors.request.use(config => {
    if (config.url === apiRoutes.refresh) {
        return config;
    }
    return config;
}, error => Promise.reject(error));

api.interceptors.response.use(
    response => response,
    async error => {
        const originalRequest = error.config;
        if (!error.response) {
            console.error('Network error or no response');
            return Promise.reject(error);
        }

        if (originalRequest.url === apiRoutes.refresh) {
            if (error.response?.status === 401) {
                window.location.href = innerRoutes.login;
            }
            return Promise.reject(error);
        }

        if (error.response?.status !== 401 || originalRequest._retry) {
            return Promise.reject(error);
        }

        if (isRefreshingFailed) {
            window.location.href = innerRoutes.login;
            return Promise.reject(error);
        }

        if (isRefreshing) {
            return new Promise((resolve, reject) => {
                failedQueue.push({ resolve, reject });
            })
                .then(() => api(originalRequest))
                .catch(err => Promise.reject(err));
        }

        originalRequest._retry = true;
        isRefreshing = true;

        try {
            const refreshResponse = await axios.post(
                apiRoutes.refresh,
                {},
                {
                    withCredentials: true,
                    headers: {
                        'Content-Type': 'application/json',
                    }
                }
            );

            if (refreshResponse.status === 200) {
            } else {
                throw new Error('Refresh failed with status: ' + refreshResponse.status);
            }

            isRefreshingFailed = false;
            processQueue(null);
            return api(originalRequest);
        } catch (refreshError) {
            console.error('Refresh failed:', refreshError);
            isRefreshingFailed = true;
            processQueue(refreshError);
            window.location.href = innerRoutes.login;
            return Promise.reject(refreshError);
        } finally {
            isRefreshing = false;
        }
    }
);

export async function get(url, params = {}, config = {}) {
    try {
        const response = await api.get(url, {
            ...config,
            params: params
        });
        return response.data;
    } catch (error) {
        console.error('GET request failed:', error.message);
        throw error;
    }
}

export async function post(url, data = {}, config = {}) {
    try {
        const response = await api.post(url, data, config);
        return response.data;
    } catch (error) {
        console.error('POST request failed:', error.message);
        throw error;
    }
}

export async function put(url, data = {}, config = {}) {
    try {
        const response = await api.put(url, data, config);
        return response.data;
    } catch (error) {
        console.error('PUT request failed:', error.message);
        throw error;
    }
}

export async function patch(url, data = {}, config = {}) {
    try {
        const response = await api.patch(url, data, config);
        return response.data;
    } catch (error) {
        console.error('PATCH request failed:', error.message);
        throw error;
    }
}

export async function del(url, config = {}) {
    try {
        const response = await api.delete(url, config);
        return response.data;
    } catch (error) {
        console.error('DELETE request failed:', error.message);
        throw error;
    }
}

export default api;