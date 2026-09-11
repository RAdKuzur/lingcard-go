import {useEffect, useState} from "react";
import { loginAxios as loginAxios } from "../../plugins/auth.js";
import { useNavigate } from "react-router-dom";
import { innerRoutes } from "../../plugins/routes.js";
import { useRedirect } from "../../hooks/useRedirect.js";
import {useAuth} from "../../plugins/AuthContext.jsx";
import {getText, lang} from "../../lang/lang.js";

export default function Login() {
    const { redirect } = useRedirect();
    const auth = useAuth()
    const navigate = useNavigate();
    const [isHover, setHover] = useState(false);
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    async function signIn() {
        setError('');
        const result = await loginAxios(name, password, auth, redirect);
        if (!result.success) {
            setError(getText(lang.login.invalidCredentials));
        }
    }

    function goToRegister() {
        redirect(innerRoutes.register);
    }

    return (
        <main className="flex flex-1 bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 items-center justify-center p-4">
            <div className="flex flex-col w-96 min-h-96 bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl shadow-indigo-500/10 text-center justify-between p-8 border border-white/50">
                <div className="font-bold text-2xl text-slate-800 mt-2">
                    {getText(lang.login.mainLabel)}
                </div>
                <div className="space-y-4">
                    <div>
                        <div className="text-sm font-medium text-slate-600 mb-1.5 text-left">{getText(lang.login.email)}</div>
                        <input
                            className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                            onInput={(e) => setName(e.target.value)}
                            value={name}
                        />
                    </div>
                    <div>
                        <div className="text-sm font-medium text-slate-600 mb-1.5 text-left">{getText(lang.login.password)}</div>
                        <input
                            className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                            type="password"
                            onInput={(e) => setPassword(e.target.value)}
                            value={password}
                        />
                    </div>
                    {error !== '' ? (
                        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-xl text-sm text-left">
                            {error}
                        </div>
                    ) : ''}
                </div>
                <div className="mt-4 space-y-3">
                    <button
                        className={`w-full py-3.5 rounded-xl bg-gradient-to-r cursor-pointer from-emerald-500 to-emerald-600 text-white font-bold transition-all duration-200 shadow-lg shadow-indigo-500/25 hover:shadow-emerald-500/40 hover:scale-[1.02] active:scale-[0.98] ${isHover ? 'from-emerald-600 to-emerald-700' : ''}`}
                        onClick={signIn}
                        onMouseEnter={() => setHover(true)}
                        onMouseLeave={() => setHover(false)}
                    >
                        {getText(lang.login.signIn)}
                    </button>
                    <div className="text-sm text-slate-600">
                        {getText(lang.login.noAccount)}{' '}
                        <button
                            onClick={goToRegister}
                            className="text-emerald-500 font-semibold hover:text-emerald-600 hover:underline transition-all duration-200 cursor-pointer"
                        >
                            {getText(lang.login.signUp)}
                        </button>
                    </div>
                </div>
            </div>
        </main>
    );
}