import {getText, lang} from "../../lang/lang.js";
import {useRedirect} from "../../hooks/useRedirect.js";
import {innerRoutes} from "../../plugins/routes.js";
import Clock from "../svg/Clock.jsx";

export default function WaitingTraining() {
    const {redirect} = useRedirect()
    function handleClick() {
        redirect(innerRoutes.profile)
    }
    return (
        <div className="flex flex-col items-center gap-8 w-96 p-10 bg-white/90 backdrop-blur-sm rounded-3xl shadow-xl border border-white/50">
            <div className="flex flex-col items-center gap-3">
                <div className="w-24 h-24 rounded-full bg-gradient-to-br from-amber-100 to-orange-100 flex items-center justify-center p-4 shadow-inner">
                   <Clock/>
                </div>
                <h2 className="text-2xl font-light text-slate-700 tracking-wide text-center">
                    <span className="text-base font-normal text-slate-500 mt-1 block">
                        {getText(lang.waitingTraining.alreadyWords)}
                        <br />
                        {getText(lang.waitingTraining.newWords)}
                    </span>
                </h2>
            </div>

            <button
                className="w-full py-3.5 px-6 bg-gradient-to-r from-amber-500 to-orange-500
                   hover:from-amber-600 hover:to-orange-600
                   text-white text-base font-medium rounded-2xl
                   transition-all duration-300 shadow-lg shadow-amber-500/25
                   hover:shadow-amber-500/40 hover:scale-[1.02] active:scale-[0.98]
                   cursor-pointer"
                onClick={handleClick}
            >
                {getText(lang.waitingTraining.continue)}
            </button>
        </div>
    );
}