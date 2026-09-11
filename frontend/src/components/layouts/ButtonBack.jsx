import {useEffect, useState} from "react";
import {useRedirect} from "../../hooks/useRedirect.js";
import {getText, lang} from "../../lang/lang.js";
import ArrowBack from "../svg/ArrowBack.jsx";

export default function ButtonBack() {
    const {redirectIfAuth} = useRedirect();
    const [hover, setHover] = useState(false);
    function goBack() {
        redirectIfAuth(-1);
    }

    return (
        <div
            className={`flex items-center gap-2 px-4 py-2 rounded-xl cursor-pointer transition-all duration-200 border border-slate-200 shadow-sm ${
                hover ? 'bg-slate-100 border-slate-300 shadow' : 'bg-white'
            }`}
            onClick={goBack}
            onMouseEnter={() => setHover(true)}
            onMouseLeave={() => setHover(false)}
        >
            <ArrowBack/>
            <span className="text-sm font-medium text-slate-600">{getText(lang.back.button)}</span>
        </div>
    );
}