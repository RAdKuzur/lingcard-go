import {useEffect, useRef, useState} from "react";
import {get, post} from "../../plugins/request.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";
import ButtonBack from "../layouts/ButtonBack.jsx";
import {getText, lang} from "../../lang/lang.js";
import {useRedirect} from "../../hooks/useRedirect.js";
import {innerRoutes} from "../../plugins/routes.js";
import ArrowDown from "../svg/ArrowDown.jsx";
import Choose from "../svg/Choose.jsx";
export default function CreateArticle() {
    const {redirect} = useRedirect()
    const [loading, setLoading] = useState(false)
    const [title, setTitle] = useState('')
    const [languageId, setLanguageId] = useState(0)
    const [content, setContent] = useState('')
    const [address, setAddress] = useState('')
    const [languages, setLanguages] = useState([])
    const [isLanguageDropdownOpen, setIsLanguageDropdownOpen] = useState(false);
    const languageDropdownRef = useRef(null);
    const [languagePost, setLanguagePost] = useState(localStorage.getItem('lang') ?? 'en')
    useEffect(() => {
        const fetchArticle = async () => {
            const response = await get(apiRoutes.languages, {}, {withCredentials: true})
            const data = await response.data
            setLanguages(data)
            const enLanguage = data.find(item => item.code === 'en');
            const enId = enLanguage?.id;
            setLanguageId(enId)
            setLanguagePost(localStorage.getItem('lang') ?? 'en')
        };
        fetchArticle();
    }, []);
    const handleLanguageChange = (value, id) => {
        setLanguageId(id)
        setLanguagePost(value);
        setIsLanguageDropdownOpen(false);
    };
    async function handleCreatePost() {
        if(address !== '' && title !== ''  && content !== '' && languageId !== 0) {
            try {
                await post(apiRoutes.posts, {
                    title: title,
                    content: content,
                    language_id: languageId,
                    address: address
                }, {withCredentials: true})

                setTitle('')
                setContent('')
                setAddress('')
            }
            catch {

            }
            setLoading(true)
            setTimeout(() => {
                setLoading(false)
                redirect(innerRoutes.home)
            }, 7000)
        }
    }
    const selectedLanguage = languages.find(l => l.code === languagePost);
    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="flex max-w-5xl mx-auto justify-start mb-6">
                <ButtonBack/>
            </div>
            <div className="max-w-5xl mx-auto space-y-8">
                <div
                    className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-8 transition-all hover:shadow-xl hover:shadow-slate-300/50">
                    <div className="flex items-center gap-3 mb-4">
                        <h1 className="text-2xl font-bold text-slate-800">{getText(lang.createArticle.create)}</h1>
                    </div>
                    <div className={'mb-6'}>
                        <div
                            className="text-sm font-medium text-slate-600 mb-1.5 text-left">{getText(lang.createArticle.title)}</div>
                        <input
                            className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                            onInput={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className={'mb-6'}>
                        <div
                            className="text-sm font-medium text-slate-600 mb-1.5 text-left">{getText(lang.createArticle.content)}</div>
                        <textarea
                            className="w-full rounded-xl h-96 px-4 py-3 border border-slate-200 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                            onInput={(e) => setContent(e.target.value)}/>
                    </div>
                    <div className={'mb-6'}>
                        <div className="text-sm font-medium text-slate-600 mb-1.5 text-left">
                            {getText(lang.createArticle.language)}
                        </div>
                        <div className="flex items-center gap-2 sm:gap-4">
                            <div className="relative" ref={languageDropdownRef}>
                                <button
                                    onClick={() => setIsLanguageDropdownOpen(!isLanguageDropdownOpen)}
                                    className="flex items-center gap-2 border border-slate-200 rounded-lg
                                                     pl-2 sm:pl-2 pr-2 sm:pr-2 py-1 sm:py-1.5
                                                     text-xs sm:text-sm font-medium text-slate-700
                                                     hover:border-indigo-400 focus:outline-none focus:ring-2
                                                     focus:ring-indigo-500/20 focus:border-indigo-500
                                                     transition-all duration-200 cursor-pointer
                                                     min-w-[60px] sm:min-w-[80px] md:min-w-[120px]
                                                     bg-transparent bg-white"
                                >
                                    <img
                                        src={`/flags/${selectedLanguage?.code}.svg`}
                                        alt={selectedLanguage?.name}
                                        className="w-4 h-4 sm:w-5 sm:h-5 rounded-sm object-cover"
                                    />
                                    <span className="hidden md:inline">{selectedLanguage?.name}</span>
                                    <ArrowDown/>
                                </button>
                                {isLanguageDropdownOpen && (
                                    <div className="absolute left-0 mt-1 min-w-[160px] bg-white
                                                      border border-slate-200 rounded-lg shadow-lg
                                                      py-1 z-50 animate-fadeIn">
                                        {languages.map((lang) => (
                                            <button
                                                key={lang.id}
                                                onClick={() => handleLanguageChange(lang.code, lang.id)}
                                                className="w-full flex items-center gap-3 px-4 py-2
                                                             text-sm text-slate-700 hover:bg-indigo-50
                                                             transition-colors duration-150"
                                            >
                                                <img
                                                    src={`/flags/${lang?.code}.svg`}
                                                    alt={lang.name}
                                                    className="w-5 h-5 rounded-sm object-cover"
                                                />
                                                <span>{lang.name}</span>
                                                {lang.id === languagePost && (
                                                    <Choose/>
                                                )}
                                            </button>
                                        ))}
                                    </div>
                                )}
                            </div>
                        </div>
                    </div>
                    <div className={'mb-6'}>
                        <div
                            className="text-sm font-medium text-slate-600 mb-1.5 text-left">{getText(lang.createArticle.address)}</div>
                        <input
                            className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                            onInput={(e) => setAddress(e.target.value)}
                        />
                    </div>
                    {
                        !loading ? (<button
                                className={'bg-indigo-500 font-bold p-2 text-white rounded-2xl cursor-pointer'}
                                onClick={handleCreatePost}>{getText(lang.createArticle.button)}
                            </button>) :
                            <div className="flex mb-6 items-center justify-center">
                                <div
                                    className="inline-flex items-center gap-2 p-2 px-6 bg-emerald-50 text-emerald-700 font-medium rounded-full border border-emerald-200 shadow-sm">
                                    {getText(lang.createArticle.article)}
                                </div>
                            </div>
                    }

                </div>

            </div>
        </main>
    );
}