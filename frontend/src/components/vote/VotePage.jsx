import {useEffect, useState} from "react";
import {get, post} from "../../plugins/request.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";
import ButtonBack from "../layouts/ButtonBack.jsx";
import {getLabel, getText, lang} from "../../lang/lang.js";
import VoiceIcon from "../svg/VoiceIcon.jsx";
import ColorChoose from "../svg/ColorChoose.jsx";
export default function VotePage() {
    const [vote, setVote] = useState([])
    const [voteOptions, setVoteOptions] = useState([])
    const [selectedOptionId, setSelectedOptionId] = useState(null);
    useEffect(() => {
        const fetchVote = async () => {
            const id = window.location.pathname.split('/').pop();
            const response = await get(apiRoutes.votes + '/' + id, {}, { withCredentials: true });
            const data = await response.data;
            setVote(data)
            setVoteOptions(data.vote_options)
            setSelectedOptionId(data.voted)
        };
        fetchVote();
    }, [])

    async function voice(id) {
        const response = await post(apiRoutes.voice + '/' + id, {}, {withCredentials: true})
    }
    async function cancelVoice(id) {
        const response = await post(apiRoutes.cancelVoice + '/' + id, {}, {withCredentials: true})
    }
    function handleVote(optionId) {
        if (selectedOptionId === optionId) {
            cancelVoice(optionId)
            setSelectedOptionId(null)
        }
        else {
            voice(optionId)
            setSelectedOptionId(optionId);
        }
    }

    return (
        <main className="min-h-screen bg-gradient-to-br via-white p-6">
            <div className="flex max-w-7xl mx-auto justify-start mb-6">
                <ButtonBack/>
            </div>
            <div className="max-w-7xl mx-auto">
                <div className="space-y-8">
                    <div className="px-4">
                        <h1 className="text-3xl md:text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-emerald-600 to-emerald-600">
                            {getLabel(vote.title)}
                        </h1>
                    </div>
                    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4">
                        {voteOptions.length > 0 ? (
                            voteOptions.map((e) => {
                                const isSelected = selectedOptionId === e.id;
                                const parsedContent = JSON.parse(e.content);
                                const currentCount =
                                    e.count +
                                    (isSelected && vote.voted !== e.id ? 1 : 0) +
                                    (!isSelected && vote.voted === e.id ? -1 : 0);

                                return (
                                    <div
                                        key={e.id}
                                        onClick={() => handleVote(e.id)}
                                        className={`
                                        group relative overflow-hidden rounded-2xl 
                                        bg-white/80 backdrop-blur-sm 
                                        border-2 transition-all duration-300 
                                        hover:shadow-xl hover:scale-[1.05] active:scale-[0.95] 
                                        cursor-pointer
                                        ${isSelected
                                            ? 'border-emerald-500 shadow-lg shadow-emerald-200/50 ring-2 ring-emerald-300/30'
                                            : 'border-slate-200/70 hover:border-emerald-300'
                                        }
                                    `}
                                    >
                                        <div className={`
                                        absolute inset-0 transition-opacity duration-300 
                                        bg-gradient-to-br from-emerald-50/0 to-emerald-50/0 
                                        ${isSelected ? 'opacity-100' : 'group-hover:opacity-100'}
                                    `}/>

                                        <div className="relative p-4 flex flex-col items-center text-center">
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
                                                        src={parsedContent.picture}
                                                        alt={parsedContent.code}
                                                        className="w-full h-full object-cover"
                                                    />
                                                </div>
                                                {isSelected && (
                                                    <div
                                                        className="absolute -top-1 -right-1 w-6 h-6 bg-emerald-500 rounded-full flex items-center justify-center shadow-lg shadow-emerald-400/50 animate-in zoom-in duration-300">
                                                        <ColorChoose></ColorChoose>
                                                    </div>
                                                )}
                                            </div>
                                            <h3 className="mt-3 text-sm font-semibold text-slate-800 leading-tight line-clamp-2">
                                                {e.title}
                                            </h3>
                                            {parsedContent.code && (
                                                <span
                                                    className="mt-1 text-[10px] font-medium text-slate-400 bg-slate-100 px-2 py-0.5 rounded-full">
                                                {parsedContent.code}
                                            </span>
                                            )}
                                            <div
                                                className="mt-3 pt-2 border-t border-slate-200/60 w-full flex items-center justify-center gap-1.5">
                                                <VoiceIcon></VoiceIcon>
                                                <span className="text-lg font-bold text-slate-700">
                                                    {currentCount}
                                                </span>
                                            </div>
                                        </div>
                                    </div>
                                );
                            })
                        ) : (
                            <></>
                        )}
                    </div>
                </div>
            </div>
        </main>
    );
}