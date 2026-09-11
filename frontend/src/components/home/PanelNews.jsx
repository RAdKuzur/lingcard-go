import {useState} from "react";
import {useRedirect} from "../../hooks/useRedirect.js";

export default function PanelNews({title, link}) {
    const [isHover, setHover] = useState(false)
    const {redirectIfAuth} = useRedirect()

    function handleClick() {
        redirectIfAuth(link)
    }
    return (
        <div
            className={`group relative h-64 bg-white rounded-2xl shadow-lg shadow-slate-200/50 hover:shadow-xl hover:shadow-slate-300/50 flex flex-col items-center justify-center cursor-pointer transition-all duration-300 ${isHover ? 'scale-105 -translate-y-1' : 'scale-100'}`}
            onMouseEnter={() => setHover(true)}
            onMouseLeave={() => setHover(false)}
            onClick={handleClick}
        >
            <div className="font-bold text-lg text-slate-700 group-hover:text-indigo-600 transition-colors duration-300">
                {title}
            </div>
        </div>
    );
}