import { useEffect, useState } from "react";
import Word from "../layouts/Word.jsx";
import ButtonBack from "../layouts/ButtonBack.jsx";
import { get } from "../../plugins/request.js";
import { apiRoutes } from "../../plugins/apiRoutes.js";
import {getText, lang} from "../../lang/lang.js";
import Loading from "../layouts/Loading.jsx";

export default function Progress() {
    const [isLoading, setLoading] = useState(false)
    const [activeTab, setActiveTab] = useState(1);
    const [page, setPage] = useState(1);
    const [limit] = useState(10);
    const [words, setWords] = useState([]);
    const [amountWords, setAmountWords] = useState(1);
    const [search, setSearch] = useState('')
    useEffect(() => { handleTabClick(1); }, []);

    function handleTabClick(tabId) {
        setActiveTab(tabId);
        setPage(1);
        handleProgress(tabId, 1, limit, search);
    }

    function nextPage() {
        if (page !== totalPages()) {
            const newPage = page + 1;
            setPage(newPage);
            handleProgress(activeTab, newPage, limit, search);
        }
    }

    function prevPage() {
        if (page > 1) {
            const newPage = page - 1;
            setPage(newPage);
            handleProgress(activeTab, newPage, limit, search);
        }
    }

    async function handleProgress(status, page = 1, limit = 10, value = '') {
        setLoading(true)
        let url = `${apiRoutes.progress}/${status}?page=${page}&limit=${limit}`;
        if (value) {
            url = url + '&search=' + value;
        }
        setPage(page);
        const response = await get(
            url,
            null,
            { withCredentials: true }
        );
        const data = await response.data;
        setAmountWords(await response.amountWords);
        setWords(data);
        setSearch(value)
        setLoading(false)
    }

    function totalPages() {
        return Math.floor((amountWords + limit - 1) / limit);
    }

    const tabs = [
        { id: 1, label: getText(lang.progress.newWords), color: 'red', onClick: () => handleTabClick(1) },
        { id: 2, label: getText(lang.progress.learningWords), color: 'blue', onClick: () => handleTabClick(2) },
        { id: 3, label: getText(lang.progress.learnedWords), color: 'green', onClick: () => handleTabClick(3) }
    ];

    const getTabStyles = (tabId, color) => {
        const isActive = activeTab === tabId;
        const baseStyles = "flex-1 px-4 py-2.5 rounded-lg font-medium transition-all duration-200 cursor-pointer";

        const activeStyles = {
            red: "bg-red-600 text-white shadow-md shadow-red-500/25",
            blue: "bg-blue-600 text-white shadow-md shadow-blue-500/25",
            green: "bg-emerald-600 text-white shadow-md shadow-emerald-500/25"
        };

        const inactiveStyles = {
            red: "bg-red-400 text-white hover:bg-red-600",
            blue: "bg-blue-400 text-white hover:bg-blue-600",
            green: "bg-emerald-400 text-white hover:bg-emerald-600"
        };

        return `${baseStyles} ${isActive ? activeStyles[color] : inactiveStyles[color]}`;
    };

    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="max-w-5xl mx-auto space-y-6">
                <div className="flex items-center gap-4">
                    <ButtonBack />
                    <h1 className="text-2xl font-bold text-slate-800">{getText(lang.progress.progressLabel)}</h1>
                </div>

                <div className="flex gap-2 p-1.5 rounded-xl shadow-lg shadow-slate-200/50">
                    {tabs.map(tab => (
                        <button
                            key={tab.id}
                            onClick={tab.onClick}
                            className={getTabStyles(tab.id, tab.color)}
                        >
                            {tab.label}
                        </button>
                    ))}
                </div>

                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-6">
                    <div className="flex items-center justify-between mb-4 gap-4 flex-wrap">
                        <h2 className="text-sm font-medium text-slate-500">{getText(lang.progress.words)} ({words.length > 0 ? amountWords : 0})</h2>
                        <div className="flex-1 max-w-xs">
                            <input
                                type="text"
                                placeholder={''}
                                className="w-full px-4 py-2 text-sm border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
                                onInput={(e) => handleProgress(activeTab, page, limit, e.target.value)}
                            />
                        </div>
                    </div>
                    {isLoading ? (<Loading/>) : (
                        <div className="space-y-3 max-h-[50vh] overflow-y-auto">
                            {words.length > 0 ? (
                                words.map((e) => (
                                    <Word
                                        key={e.id}
                                        word={e.text}
                                        translation={e.translation}
                                        level={e.level}
                                        repeat={e.repeat_time}
                                        progressId={e.id}
                                        activeTab={activeTab}
                                        transcription={e.transcription}
                                    />
                                ))
                            ) : (
                                <div className="text-center py-12 text-slate-400">
                                    <p className="text-lg">{getText(lang.progress.noWords)}</p>
                                </div>
                            )}
                        </div>
                    )}
                </div>

                {words.length > 0 && (
                    <div
                        className="flex items-center justify-between gap-4 bg-white px-4 py-3 rounded-xl shadow-lg shadow-slate-200/50">
                        <button
                            onClick={prevPage}
                            disabled={page === 1}
                            className="px-4 py-2 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-600 font-medium transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed cursor-pointer"
                        >
                            {getText(lang.progress.prev)}
                        </button>
                        <span className="text-sm text-slate-500 font-medium">
                            {page} / {totalPages()}
                        </span>
                        <button
                            onClick={nextPage}
                            disabled={page === totalPages()}
                            className="px-4 py-2 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-600 font-medium transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed cursor-pointer"
                        >
                            {getText(lang.progress.next)}
                        </button>
                    </div>
                )}
            </div>
        </main>
    );
}