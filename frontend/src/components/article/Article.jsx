import {useEffect, useState} from "react";
import {get, post, del} from "../../plugins/request.js";
import {apiCreateComment, apiRoutes} from "../../plugins/apiRoutes.js";
import ButtonBack from "../layouts/ButtonBack.jsx";
import {getText, lang} from "../../lang/lang.js";
import Location from "../svg/Location.jsx";
import View from "../svg/View.jsx";
import Like from "../svg/Like.jsx";
import Dislike from "../svg/Dislike.jsx";
import Pin from "../svg/Pin.jsx";
import Bin from "../svg/Bin.jsx";
import {useAuth} from "../../plugins/AuthContext.jsx";
export default function Article() {
    const [article, setArticle] = useState([])
    const [isLike, setLike] = useState(false)
    const [isDislike, setDislike] = useState(false)
    const [likeCount, setLikeCount] = useState(0)
    const [dislikeCount, setDislikeCount] = useState(0)
    const [comments, setComments] = useState([])
    const [textComment, setTextComment] = useState('')
    const auth = useAuth();
    useEffect(() => {
        const fetchArticle = async () => {
            const id = window.location.pathname.split('/').pop();
            const response = await get(apiRoutes.article + '/' + id, {}, { withCredentials: true });
            const data = await response.data;
            setArticle(data);
            setLike(data.is_liked)
            setDislike(data.is_disliked)
            setLikeCount(data.likes_count)
            setDislikeCount(data.dislikes_count)
            setComments(data.comments)
        };
        fetchArticle();
    }, [])


    async function like() {
        const id = window.location.pathname.split('/').pop();
        if (isDislike) {
            setDislike(false)
            setDislikeCount(prev => prev - 1)
            await reactUnset(id)
            setLike(true)
            setLikeCount(prev => prev + 1)
            await reactLike(id)
        } else if (!isLike) {
            setLike(true)
            setLikeCount(prev => prev + 1)
            await reactLike(id)
        } else {
            setLike(false)
            setLikeCount(prev => prev - 1)
            await reactUnset(id)
        }
    }

    async function dislike() {
        const id = window.location.pathname.split('/').pop();
        if (isLike) {
            setLike(false)
            setLikeCount(prev => prev - 1)
            await reactUnset(id)
            setDislike(true)
            setDislikeCount(prev => prev + 1)
            await reactDislike(id)
        } else if (!isDislike) {
            setDislike(true)
            setDislikeCount(prev => prev + 1)
            await reactDislike(id)
        } else {
            setDislike(false)
            setDislikeCount(prev => prev - 1)
            await reactUnset(id)
        }
    }

    async function reactLike(id) {
        await post(apiRoutes.like + '/' + id, {}, {withCredentials: true})
    }

    async function reactDislike(id) {
        await post(apiRoutes.dislike + '/' + id, {}, {withCredentials: true})
    }

    async function reactUnset(id) {
        await post(apiRoutes.unset + '/' + id, {}, {withCredentials: true})
    }
    async function handleSendComment(text) {
        const id = window.location.pathname.split('/').pop();
        if (text !== '') {
            await post(apiCreateComment(id), {
                text: text
            }, {withCredentials: true})
            window.location.reload()
        }
    }

    async function handleDeleteComment(id) {
        await del(apiRoutes.comments + '/' + id, {withCredentials: true})
        window.location.reload()
    }

    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="flex max-w-5xl mx-auto justify-start mb-6">
                <ButtonBack/>
            </div>
            <div className="max-w-5xl mx-auto space-y-8">
                <div key={article.id}
                     className="bg-white items-center gap-3 mb-4 shadow rounded-3xl p-8 transition-all">
                    <div className={'flex justify-between items-start'}>
                        <div className="flex items-center gap-3 mb-4">
                            <h1 className="text-xl font-bold text-slate-800">{article.title}</h1>
                        </div>
                        <div className="flex items-center gap-2 mb-4">
                            <img
                                src={`/flags/${article.code}.svg`}
                                alt={article.code}
                                className="w-6 h-6 rounded-sm object-cover"
                            />
                            <span className="text-sm font-medium text-slate-700">
                                {article.code}
                            </span>
                            <span className="text-sm text-slate-500 ml-1">
                                {article.date}
                            </span>
                        </div>
                    </div>
                    <div>
                        <p className="text-slate-600 leading-relaxed text-lg">
                            {article.content}
                        </p>
                    </div>
                    <div className="mt-4 pt-4 border-t border-slate-100 flex flex-wrap gap-4 justify-between items-center">
                        <div className="flex flex-wrap gap-4 items-center">
                            <div className="flex items-center gap-2">
                                <span className="text-sm text-slate-600 font-medium">
                                    {article.username}
                                </span>
                            </div>
                            <div className="flex items-center gap-2">
                                <Location/>
                                <span className="text-sm text-slate-500">
                                    {article.address}
                                </span>
                            </div>
                        </div>
                        <div className="flex flex-wrap gap-4 items-center">
                            <span className="text-sm text-slate-600 font-medium flex items-center gap-2">
                                <View/>
                                {article.views_count ?? 0}
                            </span>
                            <span className="text-sm text-slate-600 font-medium flex items-center gap-2">
                                <Like isLike={isLike} like={like}/>
                                {likeCount}
                            </span>
                            <span className="text-sm text-slate-600 font-medium flex items-center gap-2">
                                <Dislike isDislike={isDislike} dislike={dislike}/>
                                {dislikeCount}
                            </span>
                        </div>
                    </div>
                </div>
                <div id={"comments"}
                     className={"bg-white items-center gap-3 mb-4 shadow rounded-3xl p-8 transition-all"}>
                    <div>
                        <div className="flex items-center gap-3 mb-4">
                            <h1 className="text-xl font-bold text-slate-800">{getText(lang.article.comments)}</h1>
                        </div>
                        <div className="flex items-center gap-3 mb-4">
                            <input className={'w-full rounded-2xl p-2 outline-2 border-black'} onInput={(e) => {
                                setTextComment(e.target.value)
                            }}></input>
                            <button className={'bg-emerald-500 p-2 rounded-2xl cursor-pointer hover:bg-emerald-600'}
                                    onClick={() => handleSendComment(textComment)}>
                                <span className={'text-white font-bold'}>
                                    {getText(lang.article.send)}
                                </span>
                            </button>
                        </div>
                        <hr className="border-black-500 mb-3 mt-3"/>
                        {comments.length > 0 ? (
                            comments.map((e) => (
                                <div key={e.id}
                                     className="bg-white items-center p-5 transition-all">
                                    <div className={'flex justify-between items-start'}>
                                        <div className="flex items-center gap-3 mb-4">
                                            <h6 className="font-bold text-slate-800">{e.username}</h6>
                                            {e.is_fixed && (
                                                <Pin/>
                                            )}
                                        </div>
                                        <div className="flex items-center gap-2 mb-4">
                                            <img
                                                src={`/flags/${e.language_code}.svg`}
                                                alt={e.language_code}
                                                className="w-6 h-6 rounded-sm object-cover"
                                            />
                                            <span className="text-sm font-medium text-slate-700">
                                                {e.language_code}
                                            </span>
                                            <span className="text-sm text-slate-500 ml-1">
                                                {e.time}
                                            </span>
                                            {
                                                auth.user?.username === e.username ? (<button
                                                    className="flex items-center gap-2 p-1 rounded-xl bg-gradient-to-r from-red-500 to-rose-500 text-white text-sm font-medium shadow-sm hover:shadow-md hover:scale-105 hover:from-red-600 hover:to-rose-600 active:scale-95 transition-all duration-200 cursor-pointer"
                                                    onClick={() => handleDeleteComment(e.id)}
                                                >
                                                    <Bin/>
                                                </button>) : <></>
                                            }
                                        </div>
                                    </div>
                                    <div>
                                        <p className="text-slate-600 leading-relaxed text-lg">
                                            {e.text}
                                        </p>
                                    </div>
                                </div>
                            ))
                        ) : (
                            <div className="text-center">
                                <p className="text-slate-500 text-lg">{getText(lang.article.noComments)}</p>
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </main>
    );
}