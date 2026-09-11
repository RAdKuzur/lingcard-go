import {getText, lang} from "../../lang/lang.js";
import {useRedirect} from "../../hooks/useRedirect.js";
import {innerRoutes} from "../../plugins/routes.js";
import Success from "../svg/Success.jsx";

export default function EndTraining() {
    const {redirect} = useRedirect()
    function handleFinish() {
        redirect(innerRoutes.profile)
    }
    return (
        <div className="flex flex-col items-center gap-8 w-96 p-10 bg-white/90 backdrop-blur-sm rounded-3xl shadow-xl border border-white/50">
            <div className="flex flex-col items-center gap-3">
                <div className="w-24 h-24 rounded-full bg-gradient-to-br from-emerald-100 to-green-100 flex items-center justify-center p-4 shadow-inner">
                    <Success/>
                </div>
                <h2 className="text-2xl font-light text-slate-700 tracking-wide text-center">
                    <span className="font-bold block">{getText(lang.endTraining.congratulations)}</span>
                    <span className="text-base font-normal text-slate-500 mt-1 block">
                   {getText(lang.endTraining.message)}
                </span>
                </h2>
            </div>

            <button
                className="w-full py-3.5 px-6 bg-gradient-to-r from-emerald-500 to-green-500
                   hover:from-emerald-600 hover:to-green-600
                   text-white text-base font-medium rounded-2xl
                   transition-all duration-300 shadow-lg shadow-emerald-500/25
                   hover:shadow-emerald-500/40 hover:scale-[1.02] active:scale-[0.98]
                   cursor-pointer"
                onClick={handleFinish}
            >
                {getText(lang.endTraining.continue)}
            </button>
        </div>
    );
}