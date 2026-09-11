import {useEffect, useState} from "react";
import axios from "axios";
import { apiRoutes } from "../../plugins/apiRoutes.js";
import { getText, lang as language } from "../../lang/lang.js";
import ConfirmRegister from "./ConfirmRegister.jsx";
import ColorChoose from "../svg/ColorChoose.jsx";
import {get} from "../../plugins/request.js";
import Loading from "../layouts/Loading.jsx";

export default function Register() {
    const [step, setStep] = useState(1);
    const [lang, setLang] = useState(null);
    const [targetLang, setTargetLang] = useState(null);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const [isSuccess, setIsSuccess] = useState(false);
    const [baseLanguages, setBaseLanguages] = useState([])
    const [targetLanguages, setTargetLanguages] = useState([])
    const [selectedOptionId, setSelectedOptionId] = useState(null);
    const [selectedTargetLanguageId, setSelectedTargetLanguageId] = useState(null);

    const isStepValid = () => {
        if (step === 1) return lang !== null;
        if (step === 2) return targetLang !== null && targetLang !== lang;
        if (step === 3) return username.length >= 3 && password.length >= 6;
        return false;
    };

    useEffect(() => {
        handleBaseLanguages()
    }, []);

    async function handleBaseLanguages() {
        const response = await get(apiRoutes.languages, null, {withCredentials: true})
        const data = await response.data
        setBaseLanguages(data)
    }

    async function handleTargetLanguages(id) {
        const response = await get(apiRoutes.exceptLanguage + '/' + id, null, {withCredentials: true})
        const data = await response.data
        setTargetLanguages(data)
    }

    const nextStep = () => {
        if (isStepValid()) {
            setStep(step + 1);
            setMessage('');
        } else {
            if (step === 1) setMessage(getText(language.home.chooseBaseLang));
            if (step === 2) setMessage(getText(language.home.chooseTargetLang));
            if (step === 3) setMessage(getText(language.home.inputUsername));
        }
    };

    const prevStep = () => {
        if (step > 1) {
            setStep(step - 1);
            setMessage('');
        }
    };

    function handlePick(optionId) {
        if (selectedOptionId === optionId) {
            setLang(null)
            setSelectedOptionId(null)
        } else {
            setLang(optionId)
            handleTargetLanguages(optionId);
            setSelectedOptionId(optionId);
        }
    }

    function handleTargetPick(optionId) {
        if (selectedTargetLanguageId === optionId) {
            setTargetLang(null)
            setSelectedTargetLanguageId(null)
        } else {
            setTargetLang(optionId)
            setSelectedTargetLanguageId(optionId);
        }
    }

    async function signUp() {
        setMessage('');
        setIsSuccess(false);

        try {
            const response = await axios.post(
                apiRoutes.register,
                {
                    password: password,
                    base_language_id: lang,
                    target_language_id: targetLang,
                    name: username
                },
                {
                    withCredentials: true,
                    headers: {
                        'Content-Type': 'application/json',
                    }
                }
            );

            const successRegister = response.data.data.status;

            if (successRegister) {
                setIsSuccess(true);
            } else {
                setMessage(getText(language.register.failed));
                setIsSuccess(false);
            }
        } catch (error) {
            setMessage(getText(language.register.failed));
            setIsSuccess(false);
        }
    }

    if (isSuccess) {
        return <ConfirmRegister />;
    }

    return (
        <main className="flex flex-1 bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 items-center justify-center p-4">
            <div className="flex flex-col w-full max-w-2xl min-h-96 bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl shadow-indigo-500/10 p-8 border border-white/50">
                <div className="flex justify-center gap-2 mb-6">
                    {[1, 2, 3].map((i) => (
                        <div
                            key={i}
                            className={`h-2 rounded-full transition-all duration-300 ${
                                i === step ? 'w-12 bg-blue-600' :
                                    i < step ? 'w-8 bg-emerald-500' : 'w-8 bg-slate-200'
                            }`}
                        />
                    ))}
                </div>

                <div className="text-center mb-6">
                    <h2 className="font-bold text-2xl text-slate-800">
                        {step === 1 && getText(language.register.mainLabelStep1)}
                        {step === 2 && getText(language.register.mainLabelStep2)}
                        {step === 3 && getText(language.register.mainLabelStep3)}
                    </h2>
                    <p className="text-sm text-slate-500 mt-1">
                        {step === 1 && getText(language.register.hintLabelStep1)}
                        {step === 2 && getText(language.register.hintLabelStep2)}
                        {step === 3 && getText(language.register.hintLabelStep3)}
                    </p>
                </div>

                {message && (
                    <div className="px-4 py-3 mb-4 rounded-xl text-sm bg-red-50 border border-red-200 text-red-700">
                        {message}
                    </div>
                )}

                <div className="flex-1">
                    {step === 1 && (
                        <div>
                            <div className="text-sm font-medium text-slate-600 mb-3">
                                {getText(language.register.yourLang)}
                            </div>
                            <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
                                {baseLanguages.length > 0 ? (
                                    baseLanguages.map((e) => {
                                        const isSelected = selectedOptionId === e.id;
                                        return (
                                            <div
                                                key={e.id}
                                                onClick={() => handlePick(e.id)}
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
                                                            <div className="absolute -top-1 -right-1 w-6 h-6 bg-emerald-500 rounded-full flex items-center justify-center shadow-lg shadow-emerald-400/50">
                                                                <ColorChoose />
                                                            </div>
                                                        )}
                                                    </div>
                                                    <h3 className="mt-3 text-sm font-semibold text-slate-800 leading-tight line-clamp-2">
                                                        {e.name}
                                                    </h3>
                                                    {e.code && (
                                                        <span className="mt-1 text-[10px] font-medium text-slate-400 bg-slate-100 px-2 py-0.5 rounded-full">
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
                    )}

                    {step === 2 && (
                        <div>
                            <div className="text-sm font-medium text-slate-600 mb-3">
                                {getText(language.register.targetLang)}
                            </div>
                            <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
                                {targetLanguages.length > 0 ? (
                                    targetLanguages.map((e) => {
                                        const isSelected = selectedTargetLanguageId === e.id;
                                        return (
                                            <div
                                                key={e.id}
                                                onClick={() => handleTargetPick(e.id)}
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
                                                            <div className="absolute -top-1 -right-1 w-6 h-6 bg-emerald-500 rounded-full flex items-center justify-center shadow-lg shadow-emerald-400/50">
                                                                <ColorChoose />
                                                            </div>
                                                        )}
                                                    </div>
                                                    <h3 className="mt-3 text-sm font-semibold text-slate-800 leading-tight line-clamp-2">
                                                        {e.name}
                                                    </h3>
                                                    {e.code && (
                                                        <span className="mt-1 text-[10px] font-medium text-slate-400 bg-slate-100 px-2 py-0.5 rounded-full">
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
                    )}

                    {step === 3 && (
                        <div className="space-y-4">
                            <div>
                                <div className="text-sm font-medium text-slate-600 mb-1.5">
                                    {getText(language.register.username)}
                                </div>
                                <input
                                    className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-emerald-400 focus:ring-2 focus:ring-emerald-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                                    onInput={(e) => {
                                        const latinAndNumbers = e.target.value.replace(/[^a-zA-Z0-9]/g, '');
                                        e.target.value = latinAndNumbers;
                                        setUsername(latinAndNumbers);
                                    }}
                                    value={username}
                                />
                            </div>
                            <div>
                                <div className="text-sm font-medium text-slate-600 mb-1.5">
                                    {getText(language.register.password)}
                                </div>
                                <input
                                    className="w-full rounded-xl px-4 py-3 border border-slate-200 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-400/20 outline-none transition-all duration-200 bg-white/50 focus:bg-white"
                                    type="password"
                                    onInput={(e) => setPassword(e.target.value)}
                                />
                            </div>
                        </div>
                    )}
                </div>


                <div className="flex gap-3 mt-6">
                    {step > 1 && (
                        <button
                            onClick={prevStep}
                            className="flex-1 py-3 rounded-xl bg-slate-100 hover:bg-slate-200 text-slate-700 font-semibold transition-all duration-200 cursor-pointer hover:scale-[1.02] active:scale-[0.98]"
                        >
                            {getText(language.register.back)}
                        </button>
                    )}

                    {step < 3 ? (
                        <button
                            onClick={nextStep}
                            className={`flex-1 py-3 rounded-xl bg-gradient-to-r from-emerald-500 to-emerald-600 text-white font-bold transition-all duration-200 shadow-lg shadow-emerald-500/25 hover:shadow-emerald-500/40 cursor-pointer hover:scale-[1.02] active:scale-[0.98] ${
                                step === 1 ? 'flex-1' : ''
                            }`}
                        >
                            {getText(language.register.next)}
                        </button>
                    ) : (
                        <button
                            onClick={signUp}
                            className="flex-1 py-3 rounded-xl bg-gradient-to-r from-emerald-500 to-emerald-600 text-white font-bold transition-all duration-200 shadow-lg shadow-emerald-500/25 hover:shadow-emerald-500/40 cursor-pointer hover:scale-[1.02] active:scale-[0.98]"
                        >
                            {getText(language.register.createAccount)}
                        </button>
                    )}
                </div>
            </div>
        </main>
    );
}