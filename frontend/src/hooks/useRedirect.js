import { useNavigate } from 'react-router-dom';
import { useCallback } from 'react';
import { useAuth } from '../plugins/AuthContext.jsx';

export function useRedirect() {
    const navigate = useNavigate();
    const auth = useAuth();

    const redirect = useCallback((path, options = {}) => {
        navigate(path, options);
    }, [navigate]);

    const redirectIfAuth = useCallback((path, fallbackPath = '/login', options = {}) => {
        if (auth.isAuthenticated()) {
            navigate(path, options);
        } else {
            navigate(fallbackPath, {
                ...options,
                state: { from: path },
                replace: true
            });
        }
    }, [navigate, auth]);

    const redirectIfNotAuth = useCallback((path, fallbackPath = '/', options = {}) => {
        if (!auth.isAuthenticated()) {
            navigate(path, options);
        } else {
            navigate(fallbackPath, { replace: true });
        }
    }, [navigate, auth]);

    const goBack = useCallback(() => {
        navigate(-1);
    }, [navigate]);

    const redirectWithUserData = useCallback((path, userData = {}, options = {}) => {
        navigate(path, {
            ...options,
            state: { user: userData }
        });
    }, [navigate]);

    return {
        redirect,
        redirectIfAuth,
        redirectIfNotAuth,
        goBack,
        redirectWithUserData,
        isAuthenticated: auth.isAuthenticated,
        user: auth.user,
        getUser: auth.getUser,
        getRole: auth.getRole,
        hasRole: auth.hasRole
    };
}