import React from 'react';
import { Navigate } from 'react-router-dom';
import { innerRoutes } from "../../plugins/routes.js";
import {useAuth} from "../../plugins/AuthContext.jsx";
import {getText, lang} from "../../lang/lang.js";

const UnprotectedRoute = ({ children, redirectTo = innerRoutes.home }) => {
    const auth = useAuth();
    if (auth.isLoading) {
        return (
            <div className="loading-spinner">
                {
                    getText(lang.unprotectedRoute.loading)
                }
            </div>
        );
    }

    if (auth.isAuthenticated()) {
        return <Navigate to={redirectTo} replace />;
    }

    return children;
};

export default UnprotectedRoute;