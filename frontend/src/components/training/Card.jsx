import ButtonBack from "../layouts/ButtonBack.jsx";
import {getText, lang} from "../../lang/lang.js";
import {useEffect, useState} from "react";
import {get, patch, post} from "../../plugins/request.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";
import Cancel from "../svg/Cancel.jsx";
import Flag from "../svg/Flag.jsx";
import PaperPlane from "../svg/PaperPlane.jsx";
import {studyStatuses} from "../../plugins/studyStatus.js";

export default function Card({setTraining}) {
    const [isProblem, setProblem] = useState(false)
    const [textProblem, setTextProblem] = useState('')
    const [word, setWord] = useState(true)
    const [isHoverNo, setHoverNo] = useState(false)
    const [isHoverShow, setHoverShow] = useState(false)
    const [isHoverYes, setHoverYes] = useState(false)
    const [direction, setDirection] = useState('')
    const [opacityCard, setOpacityCard] = useState(true)
    const [opacityTranslation, setOpacityTranslation] = useState(false)
    const [loading, setLoading] = useState(true)
    const [cardId, setCardId] = useState(0)
    const [text, setText] = useState('')
    const [translation, setTranslation] = useState('')
    const [level, setLevel] = useState('')
    const [status, setStatus] = useState('')
    const [repeat, setRepeat] = useState(0)
    const [transcription, setTranscription] = useState('')
    const [isSwiping, setIsSwiping] = useState(false)

    async function trainingRepeat(status) {
        const response = await patch(apiRoutes.training + '/' + cardId, {
            status: status
        }, {withCredentials: true})
    }

    async function newWord() {
        setLoading(true)
        try {
            const response = await get(apiRoutes.training, null, {withCredentials: true});
            const data = await response.data;
            if(data) {
                setCardId(data.id)
                setText(data.text)
                setTranslation(data.translation)
                setLevel(data.level)
                setStatus(data.status)
                setRepeat(data.repeat)
                setTranscription(data.transcription)
                setWord(true)
                setOpacityTranslation(false)
                setDirection('')
                setOpacityCard(true)
                setIsSwiping(false)
            }
            else {
                setTraining(studyStatuses.learning)
            }
        } catch (error) {
            console.error('Error loading new word:', error)
        } finally {
            setLoading(false)
        }
    }

    async function handleCheckTrainingStatus() {
        try {
            const response = await get(apiRoutes.teachable, {}, {withCredentials: true})
            const data = await response.data
            setTraining(response.data.training)
            return response.data.training
        } catch (error) {
            return false
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            const status = await handleCheckTrainingStatus()
            if(status === studyStatuses.none) {
                await newWord();
            }
        }
        fetchData()
    }, []);

    function swipe(way){
        if (isSwiping) return;

        setIsSwiping(true)
        setDirection(way)
        setOpacityCard(false)

        if (way === 'left') {
            trainingRepeat(false)
        }
        if (way === 'right') {
            trainingRepeat(true)
        }

        setTimeout(async () => {
            const status = await handleCheckTrainingStatus()
            if (status === studyStatuses.none) {
                await newWord();
            } else {
                setDirection('')
                setOpacityCard(true)
                setIsSwiping(false)
                setWord(true)
                setOpacityTranslation(false)
            }
        }, 500)
    }

    function show() {
        setWord(!word);
        setOpacityTranslation(!opacityTranslation)
    }

    function handleProblem() {
        setProblem(!isProblem)
        if (isProblem) {
            setTextProblem('')
        }
    }

    async function sendProblem(problemText) {
        if (problemText) {
            try {
                const response = await post(apiRoutes.suggestions, {
                    message: JSON.stringify({
                        word: text,
                        translation: translation,
                        transcription: transcription,
                        problem: problemText
                    })
                }, {withCredentials: true})
                setTextProblem('')
                handleProblem()
            } catch (error) {
                console.error('Error sending problem:', error)
            }
        }
    }

    return (
        <div className="w-full max-w-md">
            {!loading ? (
                <>
                    <div className={`flex w-1/5 transition-all duration-500 mb-6 ${opacityCard ? 'opacity-100' : 'opacity-0  pointer-events-none'}`}>
                        <ButtonBack/>
                    </div>
                    <div className={`relative bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl shadow-indigo-500/10 p-8 transition-all duration-500 border border-white/50
                            ${direction === 'right' ? 'translate-x-full rotate-12 opacity-0 scale-90' : ''}
                            ${direction === 'left' ? '-translate-x-full -rotate-12 opacity-0 scale-90' : ''}
                            ${opacityCard ? 'opacity-100' : 'opacity-0 pointer-events-none'}
                        `}>
                        <div className="text-center">
                            <div className="flex justify-center items-center mb-6">
                                {
                                    <div className={'flex justify-start items-center w-3/5'}>
                                        {status === 1 ? (
                                            <span className="inline-flex items-center gap-1.5 px-4 py-1.5 rounded-full text-sm font-semibold bg-gradient-to-r from-emerald-400 to-emerald-500 text-white shadow-lg shadow-emerald-500/25">
                                                {getText(lang.training.newWord)}
                                            </span>
                                        ) : (
                                            <span className={`inline-flex items-center gap-1.5 px-4 py-1.5 rounded-full text-sm font-semibold bg-gradient-to-r from-amber-400 to-amber-500 text-white shadow-lg shadow-amber-500/25`}>
                                                {getText(lang.training.amountRepeat)} {repeat}
                                            </span>
                                        )}
                                    </div>
                                }
                                <div className={'flex cursor-pointer justify-end w-2/5'} onClick={handleProblem}>
                                    {!isProblem ? <Cancel/> : <Flag/>}
                                </div>
                            </div>

                            <div className="py-8">
                                <div className="text-4xl font-bold text-slate-800 mb-4 tracking-tight">
                                    {text}
                                </div>
                                {transcription !== '' && transcription !== null && (
                                    <div className={`text-2xl text-slate-600 transition-all duration-300 ${opacityTranslation ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2'}`}>
                                        [{transcription}]
                                    </div>
                                )}
                                <div className={`text-2xl text-slate-600 transition-all duration-300 ${opacityTranslation ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2'}`}>
                                    {translation}
                                </div>
                            </div>

                            {isProblem && (
                                <>
                                    <textarea
                                        className="border-2 border-black-300 focus:border-black-500 outline-none w-full rounded-2xl p-2 sm:p-4 min-h-[100px] sm:min-h-[120px] text-base"
                                        onInput={(e) => setTextProblem(e.target.value)}
                                        placeholder={getText(lang.training.problem)}
                                        value={textProblem}
                                    />
                                    <button
                                        className={`active:scale-95 p-3 rounded-2xl sm:w-auto transition-all duration-200 shadow-md hover:shadow-lg ${textProblem === '' ? 'bg-gray-500 cursor-not-allowed' : 'bg-blue-500 hover:bg-blue-600 cursor-pointer'}`}
                                        onClick={() => sendProblem(textProblem)}
                                        disabled={textProblem === ''}
                                    >
                                        <PaperPlane/>
                                    </button>
                                </>
                            )}

                            <div className="flex gap-3 mt-8">
                                <button
                                    className={`flex-1 py-3.5 rounded-xl font-semibold transition-all duration-200 shadow-lg cursor-pointer ${
                                        isHoverNo
                                            ? 'bg-red-600 shadow-red-500/40 transform scale-[1.02]'
                                            : 'bg-red-500 shadow-red-500/30 hover:bg-red-600'
                                    } text-white`}
                                    onMouseEnter={() => setHoverNo(true)}
                                    onMouseLeave={() => setHoverNo(false)}
                                    onClick={() => {
                                        swipe('left')
                                        setProblem(false)
                                    }}
                                    disabled={isSwiping}
                                >
                                    {getText(lang.training.unknown)}
                                </button>
                                <button
                                    className={`flex-1 py-3.5 rounded-xl font-semibold transition-all duration-200 shadow-lg cursor-pointer ${
                                        isHoverShow
                                            ? 'bg-blue-600 shadow-blue-500/40 transform scale-[1.02]'
                                            : 'bg-blue-500 shadow-blue-500/30 hover:bg-blue-600'
                                    } text-white`}
                                    onMouseEnter={() => setHoverShow(true)}
                                    onMouseLeave={() => setHoverShow(false)}
                                    onClick={show}
                                    disabled={isSwiping}
                                >
                                    {word ? getText(lang.training.show) : getText(lang.training.hide)}
                                </button>
                                <button
                                    className={`flex-1 py-3.5 rounded-xl font-semibold transition-all duration-200 shadow-lg cursor-pointer ${
                                        isHoverYes
                                            ? 'bg-emerald-600 shadow-emerald-500/40 transform scale-[1.02]'
                                            : 'bg-emerald-500 shadow-emerald-500/30 hover:bg-emerald-600'
                                    } text-white`}
                                    onMouseEnter={() => setHoverYes(true)}
                                    onMouseLeave={() => setHoverYes(false)}
                                    onClick={() => {
                                        swipe('right')
                                        setProblem(false)
                                    }}
                                    disabled={isSwiping}
                                >
                                    {getText(lang.training.known)}
                                </button>
                            </div>
                        </div>
                    </div>
                </>
            ) : (
                <div className="flex justify-center items-center h-64">
                    <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
                </div>
            )}
        </div>
    );
}