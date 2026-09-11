import {useRedirect} from "../../hooks/useRedirect.js";
import {innerRoutes} from "../../plugins/routes.js";
import {getText, lang} from "../../lang/lang.js";
export default function ConfirmRegister() {
    function handleLogin() {
        window.location.href = innerRoutes.news;
    }
    return (
        <main className="flex flex-1 bg-white items-center justify-center p-4">
            <div className="flex flex-col w-96 min-h-96 bg-white/90 backdrop-blur-sm rounded-3xl shadow-xl text-center justify-center items-center gap-8 p-10 border border-white/60 transition-all">
                <div
                    className="w-30 h-30 rounded-2xl flex items-center justify-center cursor-default"
                >
                    <img className={'w-full h-full'} src={'/icons/Logo.svg'} alt={'LingCard'}/>
                </div>

                <h2 className="text-2xl font-bold text-gray-800">{getText(lang.confirmRegister.welcome)}</h2>
                <p className="text-gray-500 text-sm -mt-2">{getText(lang.confirmRegister.successfulCreated)}</p>

                <button
                    onClick={handleLogin}
                    className="w-full py-4 cursor-pointer px-6 bg-gradient-to-r from-emerald-500 to-emerald-600 text-white font-semibold rounded-2xl shadow-lg shadow-emerald-500/30 hover:shadow-xl hover:shadow-emerald-500/40 hover:scale-[1.02] active:scale-95 transition-all duration-200"
                >
                    {getText(lang.confirmRegister.startTraining)}
                </button>
            </div>
        </main>
    );
}