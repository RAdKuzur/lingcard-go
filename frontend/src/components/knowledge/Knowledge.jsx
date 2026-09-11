import {useEffect, useState} from "react";
import {getText, lang as language, lang} from "../../lang/lang.js";
import LanguageKnowledge from "./LanguageKnowledge.jsx";
import ColorChoose from "../svg/ColorChoose.jsx";
import Loading from "../layouts/Loading.jsx";
import {get} from "../../plugins/request.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";

export default function Knowledge() {
    const [code, setCode] = useState(null);
    const [baseLanguages, setBaseLanguages] = useState([]);
    const [selectedOptionCode, setSelectedOptionCode] = useState(null);
    async function handleBaseLanguages() {
        const response = await get(apiRoutes.languages, null, {withCredentials: true})
        const data = await response.data
        setBaseLanguages(data)
    }

    function handlePick(langCode) {
        if (selectedOptionCode === langCode) {
            setCode(null)
            setSelectedOptionCode(null)
        } else {
            setCode(langCode)
            setSelectedOptionCode(langCode);
        }
    }

    useEffect(() => {
        handleBaseLanguages()
    }, [])
    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="max-w-5xl mx-auto space-y-6">
                <div className="flex items-center gap-4">
                    <h1 className="text-2xl font-bold text-slate-800">{getText(lang.knowledge.label)}</h1>
                </div>
            </div>
            <LanguageKnowledge code={code}/>
            <div className="max-w-5xl mx-auto space-y-6 p-8 bg-white mt-5 rounded-2xl">
                <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
                    {baseLanguages.length > 0 ? (
                        baseLanguages.map((e) => {
                            const isSelected = selectedOptionCode === e.code;
                            return (
                                <div
                                    key={e.id}
                                    onClick={() => handlePick(e.code)}
                                    className={`
                                                    group relative overflow-hidden rounded-2xl 
                                                    bg-white/80 backdrop-blur-sm 
                                                    border-2 transition-all duration-300 
                                                    hover:shadow-xl hover:scale-[1.03] active:scale-[0.97] 
                                                    cursor-pointer p-4
                                                    ${isSelected
                                        ? 'border-emerald-500 shadow-lg shadow-emerald-200/50 ring-2 ring-emerald-300/30'
                                        : 'border-slate-200/70 hover:border-emerald-300'
                                    }
                                                `}
                                >
                                    <div className="flex flex-col items-center text-center">
                                        <div className="relative">
                                            <div className={`
                                                            w-14 h-14 rounded-full overflow-hidden 
                                                            border-3 transition-all duration-300
                                                            ${isSelected
                                                ? 'border-emerald-500 shadow-lg shadow-emerald-300/50'
                                                : 'border-slate-200 group-hover:border-emerald-300'
                                            }
                                                        `}>
                                                <img
                                                    src={`/flags/${e.code}.svg`}
                                                    alt={e.code}
                                                    className="w-full h-full object-cover"
                                                />
                                            </div>
                                            {isSelected && (
                                                <div
                                                    className="absolute -top-1 -right-1 w-6 h-6 bg-emerald-500 rounded-full flex items-center justify-center shadow-lg shadow-emerald-400/50">
                                                    <ColorChoose/>
                                                </div>
                                            )}
                                        </div>
                                        <h3 className="mt-3 text-sm font-semibold text-slate-800 leading-tight line-clamp-2">
                                            {e.name}
                                        </h3>
                                        {e.code && (
                                            <span
                                                className="mt-1 text-[10px] font-medium text-slate-400 bg-slate-100 px-2 py-0.5 rounded-full">
                                                {e.code}
                                            </span>
                                        )}
                                    </div>
                                </div>
                            );
                        })
                    ) : (
                        <Loading></Loading>
                    )}
                </div>
            </div>
        </main>
    )
}