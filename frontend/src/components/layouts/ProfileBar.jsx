import { innerRoutes } from "../../plugins/routes.js";
import { useRedirect } from "../../hooks/useRedirect.js";
import { useAuth } from "../../plugins/AuthContext.jsx";
import {useEffect, useState} from "react";
import {getText, lang} from "../../lang/lang.js";

export default function ProfileBar() {
    const { redirectIfAuth } = useRedirect();
    const auth = useAuth();
    function goProfile() {
        if (auth.isAuthenticated()) {
            redirectIfAuth(innerRoutes.profile);
        } else {
            redirectIfAuth(innerRoutes.login, innerRoutes.login);
        }
    }


    const username = auth.isAuthenticated() ? auth.user?.username || getText(lang.profileBar.guest) : getText(lang.profileBar.guest);
    const role = auth.isAuthenticated() ? auth.user?.role || '' : '';

    const roleColors = {
        admin: 'bg-purple-100 text-purple-700',
        user: 'bg-blue-100 text-blue-700',
        moderator: 'bg-orange-100 text-orange-700',
    };

    const roleClass = role && roleColors[role.toLowerCase()]
        ? roleColors[role.toLowerCase()]
        : 'bg-slate-100 text-slate-700';

    return (
        <>
            <div
                className="flex items-center gap-3 px-3 py-2 rounded-xl bg-white/50 hover:bg-white border border-slate-200/50 hover:border-slate-300 shadow-sm hover:shadow-md transition-all duration-200 cursor-pointer"
                onClick={goProfile}
            >
                {auth.isAuthenticated() ? (<div
                    className="w-10 h-10 rounded-xl bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center text-white font-semibold text-sm shadow-md shadow-blue-500/20">
                    {auth.isAuthenticated() ? username.charAt(0).toUpperCase() : '👤'}
                </div>) : ''}
                <div className="flex flex-col items-start">
                <span className="text-sm font-semibold text-slate-700">
                    {auth.isAuthenticated() ? username : getText(lang.profileBar.signIn)}
                </span>
                </div>
            </div>
        </>
    );
}