import SelectLanguage from "../layouts/SelectLanguage.jsx";
import {useEffect, useState} from "react";
import ButtonBack from "../layouts/ButtonBack.jsx";
import { logoutAxios } from "../../plugins/auth.js";
import { patch, get, del } from "./../../plugins/request.js";
import { apiRoutes } from "../../plugins/apiRoutes.js";
import {useRedirect} from "../../hooks/useRedirect.js";
import {useAuth} from "../../plugins/AuthContext.jsx";
import {getText, lang} from "../../lang/lang.js";

export default function Profile() {
    const { redirectIfAuth } = useRedirect();
    const auth = useAuth()

    const [isHover1, setHover1] = useState(false);
    const [isHover2, setHover2] = useState(false);
    const [baseLang, setBaseLang] = useState(0);
    const [targetLang, setTargetLang] = useState(0);
    const [noneWords, setNoneWords] = useState(0);
    const [learningWords, setLearningWords] = useState(0);
    const [learnedWords, setLearnedWords] = useState(0);

    const [initialBaseLang, setInitialBaseLang] = useState(0);
    const [initialTargetLang, setInitialTargetLang] = useState(0);
    const [errorMessage, setErrorMessage] = useState('');
    const [successMessage, setSuccessMessage] = useState('');

    function logout() {
        logoutAxios(auth);
    }

    const validateChanges = () => {
        if (String(baseLang) === String(targetLang)) {
            setErrorMessage(getText(lang.profile.error.sameLanguages))
            return false
        }

        if (baseLang === 0 || targetLang === 0 ||
            baseLang === '0' || targetLang === '0') {
            setErrorMessage(getText(lang.profile.error.bothLanguages))
            return false
        }

        if (String(baseLang) === String(initialBaseLang) &&
            String(targetLang) === String(initialTargetLang)) {
            setErrorMessage(getText(lang.profile.error.notChanged))
            return false
        }

        return true
    }

    async function changeProfile() {
        setErrorMessage('')
        setSuccessMessage('')

        if (!validateChanges()) {
            setTimeout(() => setErrorMessage(''), 3000)
            return
        }

        try {
            const response = await patch(apiRoutes.profile, {
                base_language_id: Number(baseLang),
                target_language_id: Number(targetLang)
            }, {
                withCredentials: true
            });

            setInitialBaseLang(baseLang)
            setInitialTargetLang(targetLang)
            setSuccessMessage(getText(lang.profile.success.successChanged))

            setTimeout(() => setSuccessMessage(''), 3000)

        } catch (error) {
            setErrorMessage(getText(lang.profile.error.failedUpdate))
            setTimeout(() => setErrorMessage(''), 3000)
        }
    }

    async function clearProgress() {
        const response = await del(apiRoutes.progress, {
            withCredentials: true
        });
        window.location.reload()
    }

    useEffect(() => {
        const fetchProfile = async () => {
            try {
                const response = await get(apiRoutes.profile, {}, {withCredentials: true});
                const data = response.data;
                setTargetLang(data.target_language_id)
                setBaseLang(data.base_language_id)
                setInitialBaseLang(data.base_language_id)
                setInitialTargetLang(data.target_language_id)
                setNoneWords(data.none_words)
                setLearningWords(data.learning_words)
                setLearnedWords(data.learned_words)
            } catch (error) {

            }
        };
        fetchProfile();
    }, []);

    const isButtonDisabled = () => {
        const hasChanges = String(baseLang) !== String(initialBaseLang) ||
            String(targetLang) !== String(initialTargetLang)
        const isValid = baseLang !== 0 && targetLang !== 0 &&
            String(baseLang) !== String(targetLang)
        return !(hasChanges && isValid)
    }

    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 flex justify-center p-6">
            <div className="flex flex-col w-full max-w-5xl items-center space-y-6">
                <div className="flex w-full items-center gap-4">
                    <ButtonBack />
                    <h1 className="text-2xl font-bold text-slate-800">{getText(lang.profile.profileLabel)}</h1>
                </div>

                <div className="w-full bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-6">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                            <label className="text-sm font-medium text-slate-600 block mb-2">{getText(lang.profile.baseLang)}</label>
                            <SelectLanguage
                                setLang={setBaseLang}
                                value={baseLang}
                                onChange={() => {
                                    setErrorMessage('')
                                    setSuccessMessage('')
                                }}
                                exceptId={targetLang}
                            />
                        </div>
                        <div>
                            <label className="text-sm font-medium text-slate-600 block mb-2">{getText(lang.profile.targetLang)}</label>
                            <SelectLanguage
                                setLang={setTargetLang}
                                value={targetLang}
                                onChange={() => {
                                    setErrorMessage('')
                                    setSuccessMessage('')
                                }}
                                exceptId={baseLang}
                            />
                        </div>
                    </div>

                    {errorMessage && (
                        <div className="mt-4 p-3 bg-rose-50 border border-rose-200 text-rose-700 rounded-xl text-sm">
                            {errorMessage}
                        </div>
                    )}
                    {successMessage && (
                        <div className="mt-4 p-3 bg-emerald-50 border border-emerald-200 text-emerald-700 rounded-xl text-sm">
                            {successMessage}
                        </div>
                    )}

                    <button
                        className={'mt-4 w-full px-6 py-2.5 font-medium rounded-xl transition-all duration-200 shadow-lg ' +
                            'bg-emerald-500 hover:bg-emerald-600 text-white shadow-emerald-500/25 hover:shadow-emerald-500/35'}
                        onClick={changeProfile}
                    >
                        {getText(lang.profile.buttonChange)}
                    </button>
                </div>

                <div className="w-full bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-6">
                    <h2 className="text-sm font-medium text-slate-500 mb-4">{getText(lang.profile.stats)}</h2>
                    <div className="space-y-3">
                        <div className="flex items-center justify-between px-4 py-3 bg-slate-50 rounded-xl">
                            <span className="text-slate-700 font-medium">{getText(lang.profile.alreadyLearned)}</span>
                            <span className="px-4 py-1 rounded-full bg-emerald-500 text-white font-bold text-sm">
                                {learnedWords}
                            </span>
                        </div>
                        <div className="flex items-center justify-between px-4 py-3 bg-slate-50 rounded-xl">
                            <span className="text-slate-700 font-medium">{getText(lang.profile.learning)}</span>
                            <span className="px-4 py-1 rounded-full bg-blue-500 text-white font-bold text-sm">
                                {learningWords}
                            </span>
                        </div>
                        <div className="flex items-center justify-between px-4 py-3 bg-slate-50 rounded-xl">
                            <span className="text-slate-700 font-medium">{getText(lang.profile.noLearning)}</span>
                            <span className="px-4 py-1 rounded-full bg-rose-500 text-white font-bold text-sm">
                                {noneWords}
                            </span>
                        </div>
                    </div>
                </div>

                <div className="flex w-full gap-4">
                    <button
                        className={`flex-1 px-6 py-3 rounded-xl font-medium transition-all duration-200 border-2 cursor-pointer ${
                            isHover1
                                ? 'bg-rose-500 border-rose-500 text-white'
                                : 'bg-white border-rose-500 text-rose-500 hover:bg-rose-50'
                        }`}
                        onMouseEnter={() => setHover1(true)}
                        onMouseLeave={() => setHover1(false)}
                        onClick={logout}
                    >
                        {getText(lang.profile.signOut)}
                    </button>
                    <button
                        className={`flex-1 px-6 py-3 rounded-xl font-medium transition-all duration-200 cursor-pointer ${
                            isHover2
                                ? 'bg-rose-700 shadow-lg shadow-rose-500/35'
                                : 'bg-rose-500 shadow-lg shadow-rose-500/25 hover:bg-rose-600'
                        } text-white`}
                        onMouseEnter={() => setHover2(true)}
                        onMouseLeave={() => setHover2(false)}
                        onClick={clearProgress}
                    >
                        {getText(lang.profile.clearProgress)}
                    </button>
                </div>
            </div>
        </main>
    );
}