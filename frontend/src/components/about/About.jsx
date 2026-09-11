import ButtonBack from "../layouts/ButtonBack.jsx";
import {getText, lang} from "../../lang/lang.js";
import {useEffect, useState} from "react";
import {get, post} from "../../plugins/request.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";
import ButtonLink from "../layouts/ButtonLink.jsx";
import {donations} from "../../plugins/donation.js";

export default function About() {
    const [map, setMap] = useState([])
    const [isHidden, setHidden] = useState(true)
    const [input, setInput] = useState('')
    const [isMessage, setMessage] = useState(false)
    useEffect(() => {
        async function Languages() {
            const response = await get(apiRoutes.langMap, {}, {withCredentials: true})
            const data = await response.data
            setMap(data.map)
        }
        Languages()
    }, [])

    function handleFeedback() {
        setHidden(false)
    }
    async function sendFeedback() {
        await post(apiRoutes.suggestions, {
            message: input
        }, {withCredentials: true})
        setMessage(true)
        setTimeout(() => {
            setMessage(false)
        }, 5000)
        setHidden(true)
    }
    function hideFeedback() {
        setHidden(true)
    }
    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-4 sm:p-6">
            <div className="flex max-w-5xl mx-auto justify-start mb-4 sm:mb-6">
                <ButtonBack />
            </div>
            <div className="max-w-5xl mx-auto space-y-4 sm:space-y-6">
                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-4 sm:p-6 md:p-8 transition-all hover:shadow-xl hover:shadow-slate-300/50">
                    <div className="flex items-center gap-3 mb-3 sm:mb-4">
                        <h1 className="text-xl sm:text-2xl font-bold text-slate-800">
                            {getText(lang.about.me)}
                        </h1>
                    </div>
                    <p className="text-slate-600 leading-relaxed text-base sm:text-lg">
                        {getText(lang.about.mission)}
                    </p>
                </div>

                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-4 sm:p-6 md:p-8 transition-all hover:shadow-xl hover:shadow-slate-300/50">
                    <div className="flex items-center gap-3 mb-4">
                        <h1 className="text-xl sm:text-2xl font-bold text-slate-800">
                            {getText(lang.about.contacts)}
                        </h1>
                    </div>
                    <div className="space-y-3">
                        <a
                            href="mailto:drive16052003@gmail.com"
                            className="flex items-center gap-3 text-slate-600 hover:text-blue-600 transition-colors group break-all"
                        >
                        <span className="text-base sm:text-lg group-hover:underline">
                            drive16052003@gmail.com
                        </span>
                        </a>
                        <ButtonLink link={'https://www.reddit.com/r/lingcard/'} color={'white'} label={'Reddit'} icon={'/icons/Reddit.svg'}></ButtonLink>
                    </div>
                    <div className="mt-4">
                        {!isHidden && (
                            <div>
                            <textarea
                                className="border-2 border-black-300 focus:border-black-500 outline-none w-full rounded-2xl p-3 sm:p-4 min-h-[100px] sm:min-h-[120px] text-base"
                                onInput={(e) => setInput(e.target.value)}
                            />
                            </div>
                        )}

                        {isMessage && (
                            <span className="font-bold text-green-600 block mt-2">
                            {getText(lang.about.success)}
                        </span>
                        )}

                        <div className="flex flex-col sm:flex-row gap-3 mt-3">
                            {isHidden ? (
                                <button
                                    className="bg-emerald-500 cursor-pointer p-3 rounded-2xl w-full sm:w-auto hover:bg-emerald-600 transition-colors"
                                    onClick={handleFeedback}
                                >
                                <span className="font-bold text-white text-sm sm:text-base">
                                    {getText(lang.about.feedback)}
                                </span>
                                </button>
                            ) : (
                                <>
                                    <button
                                        className="bg-emerald-500 cursor-pointer p-3 rounded-2xl w-full sm:w-auto hover:bg-emerald-600 transition-colors"
                                        onClick={(e) => sendFeedback(e.target.value)}
                                    >
                                    <span className="font-bold text-white text-sm sm:text-base">
                                        {getText(lang.about.sendFeedback)}
                                    </span>
                                    </button>
                                    <button
                                        className="bg-red-500 cursor-pointer p-3 rounded-2xl w-full sm:w-auto hover:bg-red-600 transition-colors"
                                        onClick={hideFeedback}
                                    >
                                    <span className="font-bold text-white text-sm sm:text-base">
                                        {getText(lang.about.hide)}
                                    </span>
                                    </button>
                                </>
                            )}
                        </div>
                    </div>
                </div>
                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 overflow-hidden">
                    <div className="p-4 sm:p-6">
                        <h1 className="text-xl sm:text-2xl font-bold text-slate-800 text-center mb-4">
                            {getText(lang.about.languages)}
                        </h1>
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
                                        <span className="text-sm sm:text-base">
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