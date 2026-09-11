import {useEffect} from "react";
import {getText, lang} from "../../lang/lang.js";
import {apiRoutes} from "../../plugins/apiRoutes.js";
import axios from 'axios';
export default function LanguageKnowledge({code = null}) {
    async function download(langCode) {
        try {
            const response = await axios({
                url: apiRoutes.packages + '/' + langCode,
                method: 'GET',
                withCredentials: true,
                responseType: 'blob',
            });
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            link.download = `${langCode}.jsonl`;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            window.URL.revokeObjectURL(url);
        } catch (error) {
            console.error('Download error:', error);
        }
    }
    useEffect(() => {

    }, [])
    return (
        <>
            { !code ? (
                <></>
            ) : (
                <div className="max-w-5xl mx-auto space-y-6 p-8 bg-white mt-5 rounded-2xl">
                    <button className={`p-2 cursor-pointer rounded-2xl font-bold border-black-200 outline-2`} onClick={() => {download(code)}}>
                        <div className={`flex gap-1`}>
                            <img src={'/icons/JSON.svg'} alt={''} className="w-6 h-6 object-contain"/>
                            <span>{getText(lang.knowledge.downloadJSON)}</span>
                        </div>
                    </button>
                </div>
            )}
        </>

    )
}