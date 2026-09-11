import Word from "../layouts/Word.jsx";
import SelectLanguage from "../layouts/SelectLanguage.jsx";
import ButtonBack from "../layouts/ButtonBack.jsx";
import { useState } from "react";
import { get } from "../../plugins/request.js";
import { apiDictionary } from "../../plugins/apiRoutes.js";
import {getText, lang as language} from "../../lang/lang.js";

export default function Dictionary() {
    const [lang1, setLang1] = useState(1);
    const [lang2, setLang2] = useState(1);
    const [words, setWords] = useState([])
    const [page, setPage] = useState(1)
    const [limit, setLimit] = useState(10)
    const [amountWords, setAmountWords] = useState(1);
    const [isPaginator, setPaginator] = useState(false)
    const [search, setSearch] = useState('')
    async function handleSearch(value, page = 1, limit = 10) {

        if (lang1 !== 0 && lang2 !== 0 && lang1 !== lang2) {
            const response = await get(apiDictionary(lang1, lang2, page, limit, value), null, {withCredentials: true})
            const data = await response.data;
            setPaginator(true)
            setAmountWords(await response.amountWords)
            setWords(data)
        }
        setSearch(value)
    }
    function nextPage() {
        if(page !== totalPages()) {
            const newPage = page + 1;
            setPage(newPage);
            handleSearch(search, newPage, limit)
        }
    }

    function prevPage() {
        if(page > 1) {
            const newPage = page - 1;
            setPage(newPage)
            handleSearch(search, newPage, limit)
        }
    }

    function totalPages() {
        return Math.floor((amountWords + limit - 1) / limit);
    }

    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="max-w-5xl mx-auto space-y-6">
                <div className="flex items-center gap-4">
                    <ButtonBack />
                    <h1 className="text-2xl font-bold text-slate-800">{getText(language.dictionary.dictionary)}</h1>
                </div>

                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-6">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                            <label className="text-sm font-medium text-slate-600 block mb-2">{getText(language.dictionary.lang1)}</label>
                            <SelectLanguage setLang={setLang1} value={lang1} />
                        </div>
                        <div>
                            <label className="text-sm font-medium text-slate-600 block mb-2">{getText(language.dictionary.lang2)}</label>
                            <SelectLanguage setLang={setLang2} value={lang2}/>
                        </div>
                    </div>
                    <button
                        onClick={() => handleSearch(search, page, limit)}
                        className="mt-4 w-full px-6 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-medium rounded-xl transition-all duration-200 shadow-lg shadow-emerald-500/25 hover:shadow-emerald-500/35 cursor-pointer"
                    >
                        {getText(language.dictionary.show)}
                    </button>
                </div>

                <div className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-6">
                    <div className="flex items-center justify-between mb-4 gap-4 flex-wrap">
                        <h2 className="text-sm font-medium text-slate-500">{getText(language.dictionary.words)} ({words.length > 0 ? amountWords : 0})</h2>
                        <div className="flex-1 max-w-xs">
                            <input
                                type="text"
                                placeholder={getText(language.dictionary.searchPlaceholder)}
                                className="w-full px-4 py-2 text-sm border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
                                onInput={(e) => handleSearch(e.target.value, page, limit)}
                            />
                        </div>
                    </div>
                    <div className="space-y-3 max-h-[50vh] overflow-y-auto">
                        {words.length > 0 ? (
                            words.map((e) => (
                                <Word
                                    key={e.id}
                                    word={e.text}
                                    translation={e.translation}
                                    level={e.level}
                                    transcription={e.transcription}
                                />
                            ))
                        ) : (
                            <div className="text-center py-12 text-slate-400">
                                <p className="text-lg">{getText(language.dictionary.chooseLabel)}</p>
                            </div>
                        )}
                    </div>
                </div>

                {isPaginator && words.length > 0 && (
                    <div className="flex items-center justify-between gap-4 bg-white px-4 py-3 rounded-xl shadow-lg shadow-slate-200/50">
                        <button
                            onClick={prevPage}
                            disabled={page === 1}
                            className="px-4 py-2 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-600 font-medium transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed"
                        >
                            {getText(language.dictionary.prev)}
                        </button>
                        <span className="text-sm text-slate-500 font-medium">
                            {page} / {totalPages()}
                        </span>
                        <button
                            onClick={nextPage}
                            disabled={page === totalPages()}
                            className="px-4 py-2 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-600 font-medium transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed"
                        >
                            {getText(language.dictionary.next)}
                        </button>
                    </div>
                )}
            </div>
        </main>
    );
}