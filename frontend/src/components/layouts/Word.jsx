import {del} from "../../plugins/request.js";
import {apiClearWordProgress} from "../../plugins/apiRoutes.js";
import {useState} from "react";
import {getText, lang} from "../../lang/lang.js";
import Bin from "../svg/Bin.jsx";

export default function Word({ word, translation, level = null, repeat = null , progressId = null, activeTab = 1, transcription = ''}) {
    const [isHidden, setHidden] = useState(false)
    async function handleProgress(id) {
        const response = await del(
            apiClearWordProgress(id),
            null,
            {withCredentials: true}
        );
        setHidden(true)

    }
    function setLevelColor(level){
        switch (level) {
            case 'Начальный':
                return 'bg-green-500 hover:bg-green-600 text-white';
            case 'Базовый':
                return 'bg-lime-500 hover:bg-lime-600 text-white';
            case 'Средний':
                return 'bg-yellow-500 hover:bg-yellow-600 text-white';
            case 'Выше среднего':
                return 'bg-orange-500 hover:bg-orange-600 text-white';
            case 'Продвинутый':
                return 'bg-red-400 hover:bg-red-500 text-white';
            case 'Профессиональный':
                return 'bg-red-600 hover:bg-red-700 text-white';
            default:
                return 'bg-gray-500 hover:bg-gray-600 text-white';
        }
    }
    return (
        <div
            className={`flex items-center gap-4 w-full px-5 py-3 bg-white rounded-xl border border-slate-200 hover:border-slate-300 hover:shadow-md transition-all duration-200 ${isHidden ? 'hidden' : ''}`}>
            <div className="flex-1 min-w-0">
                <div className="font-semibold text-slate-800 truncate">{word}</div>
                {transcription === null || transcription === '' ? '' : (<div className="text-slate-500 text-sm truncate">[{transcription}]</div>)}
                <div className="text-slate-500 text-sm truncate">{translation}</div>
            </div>
            <div className="flex items-center gap-2 flex-shrink-0">
                {level !== null && (
                    <span className={`px-3 py-3 rounded-full text-xs font-medium whitespace-nowrap ${setLevelColor(level)}`}>
                    </span>
                )}
            </div>
            <div className={`flex items-center gap-2 flex-shrink-0 ${activeTab !== 1 && progressId ? '' : 'hidden'}`}>
                <button
                    className="flex items-center gap-2 px-4 py-2 rounded-xl bg-gradient-to-r from-red-500 to-rose-500 text-white text-sm font-medium shadow-sm hover:shadow-md hover:scale-105 hover:from-red-600 hover:to-rose-600 active:scale-95 transition-all duration-200 cursor-pointer"
                    onClick={() => handleProgress(progressId)}
                >
                    <Bin/>
                    <span>{getText(lang.word.clearProgress)}</span>
                </button>
            </div>
        </div>
    );
}