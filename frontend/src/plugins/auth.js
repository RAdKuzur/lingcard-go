// services/auth.js
import axios from "axios";
import { apiRoutes } from "./apiRoutes.js";
import { post } from "./request.js";
import {innerRoutes} from "./routes.js";

export async function loginAxios(name, password, authContext, redirect) {
    try {
        const response = await axios.post(
            apiRoutes.login,
            {
                name: name,
                password: password
            },
            {
                withCredentials: true
            }
        );

        const data = await response.data;

        authContext.login({
            username: data.user.username,
            role: data.user.role,
            ...data.user
        });

        redirect(innerRoutes.home);
        return { success: true, data };
    } catch (error) {
        console.error('Login failed:', error);
        return { success: false, error: error.response?.data || error.message };
    }
}

export async function logoutAxios(authContext) {
    try {
        await post(
            apiRoutes.logout,
            null,
            {
                withCredentials: true
            }
        );
    } catch (error) {
        console.error('Logout error:', error);
    } finally {
        authContext.logout();
    }
}
export function checkAuth(navigate, authContext) {
    if (authContext.isAuthenticated()) {
        return true;
    }
    navigate(innerRoutes.login);
    return false;
}