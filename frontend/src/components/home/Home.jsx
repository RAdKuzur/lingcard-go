import { useState, useEffect } from 'react';
import {apiRoutes} from "../../plugins/apiRoutes.js";
import {get} from "../../plugins/request.js";
import {useRedirect} from "../../hooks/useRedirect.js";
import {innerRoutes} from "../../plugins/routes.js";
import {useAuth} from "../../plugins/AuthContext.jsx";
import {getText, lang} from "../../lang/lang.js";

export default function Home() {
    const [map, setMap] = useState([]);
    const {redirect} = useRedirect();
    const auth = useAuth();
    useEffect(() => {
        async function Languages() {
            const response = await get(apiRoutes.langMap, {}, { withCredentials: true });
            const data = await response.data;
            setMap(data.map);
        }
        Languages();
    }, []);

    function goLogin() {
        return redirect(innerRoutes.login)
    }
    function goRegister() {
        return redirect(innerRoutes.register)
    }
    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="max-w-6xl mx-auto">
                <div className="text-center mb-8">
                    <h1 className="text-4xl sm:text-5xl font-bold text-slate-800 mb-4">
                        LingCard
                    </h1>
                    <p className="text-lg sm:text-xl text-slate-600 max-w-3xl mx-auto mb-6">
                        {getText(lang.home.mainLabel)}
                    </p>
                    {!auth.isAuthenticated() ? (
                        <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
                            <button className="cursor-pointer px-8 py-3 bg-emerald-500 hover:bg-emerald-600 text-white font-semibold rounded-lg transition-all duration-200 shadow-md hover:shadow-lg transform hover:scale-105"
                                    onClick={goRegister}
                            >
                                {getText(lang.home.startTraining)}
                            </button>
                            <button className="cursor-pointer px-8 py-3 bg-white hover:bg-slate-50 text-slate-700 font-semibold rounded-lg transition-all duration-200 shadow-md hover:shadow-lg border border-slate-200 transform hover:scale-105"
                                    onClick={goLogin}
                            >
                                {getText(lang.home.login)}
                            </button>
                        </div>)
                        :
                    <></>}
                </div>
                <div className="bg-white rounded-xl shadow-sm overflow-hidden">
                    <div className="px-4 sm:px-6 py-4 bg-slate-50 border-b border-slate-200">
                        <h2 className="text-lg font-semibold text-slate-700">
                            {getText(lang.home.availableLang)}
                        </h2>
                    </div>
                    <div className="divide-y divide-slate-100">
                        {map.map((language) => (
                            <div
                                key={language.code}
                                className="p-4 sm:p-6 hover:bg-slate-50 transition-colors"
                            >
                                <div className="flex flex-col sm:flex-row sm:items-center gap-3 sm:gap-4">
                                    <div className="flex items-center gap-3 sm:w-1/3">
                                        <img
                                            src={`/flags/${language.code}.svg`}
                                            alt={language.code}
                                            className="w-10 h-10 sm:w-12 sm:h-12 rounded-full object-cover border-2 border-slate-200 flex-shrink-0"
                                        />
                                        <span className="text-sm sm:text-base font-medium text-slate-700">
                                            {language.label}
                                        </span>
                                    </div>
                                    <div className="flex flex-wrap items-center gap-2 sm:gap-3 sm:flex-1">
                                        {language.available_codes.map((code) => (
                                            <img
                                                key={code}
                                                src={`/flags/${code}.svg`}
                                                alt={code}
                                                className="w-8 h-8 sm:w-10 sm:h-10 rounded-full object-cover border-2 border-slate-200 flex-shrink-0 hover:scale-110 transition-transform"
                                            />
                                        ))}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </main>
    );
}