import React from 'react';
import { Navigate } from 'react-router-dom';
import { innerRoutes } from "../../plugins/routes.js";
import { useAuth } from "../../plugins/AuthContext.jsx";
import {getText, lang} from "../../lang/lang.js";

const ProtectedRoute = ({ children, requiredRole = null }) => {
    const auth = useAuth();

    if (auth.isLoading) {
        return (
            <div className="flex justify-center items-center min-h-[200px]">
                <div className="text-gray-500">{getText(lang.protectedRoute.loading)}</div>
            </div>
        );
    }

    const isAuth = auth.isAuthenticated();
    if (!isAuth) {
        auth.logout();
        return <Navigate to={innerRoutes.home} replace />;
    }
    return children;
};

export default ProtectedRoute;