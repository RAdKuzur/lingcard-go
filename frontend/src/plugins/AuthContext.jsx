// context/AuthContext.js
import React, { createContext, useContext, useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import {apiRoutes} from "./apiRoutes.js";
import {post} from "./request.js";
import {innerRoutes} from "./routes.js";

const AuthContext = createContext(null);

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        checkInitialAuth();
    }, []);

    const checkInitialAuth = useCallback(async () => {
        try {
            const response = await post(apiRoutes.user, {}, { withCredentials: true });
            if (response && response.data) {
                setUser({
                    username: response.data.username,
                    role: response.data.role,
                    ...response.data
                });
            }
        } catch (error) {
            setUser(null);
        } finally {
            setIsLoading(false);
        }
    }, []);

    const login = useCallback((userData) => {
        setUser({
            username: userData.username,
            role: userData.role,
            ...userData
        });
    }, []);

    const logout = useCallback(() => {
        navigate(innerRoutes.login)
        setUser(null);
    }, []);

    const isAuthenticated = useCallback(() => {
        return user !== null;
    }, [user]);

    const getRole = useCallback(() => {
        return user?.role || null;
    }, [user]);

    const getUsername = useCallback(() => {
        return user?.username || null;
    }, [user]);

    const hasRole = useCallback((requiredRole) => {
        if (!user) return false;
        if (Array.isArray(requiredRole)) {
            return requiredRole.includes(user.role);
        }
        return user.role === requiredRole;
    }, [user]);

    const value = {
        user,
        login,
        logout,
        isAuthenticated,
        getRole,
        getUsername,
        hasRole,
        isLoading
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within AuthProvider');
    }
    return context;
}