import React, { useState, useEffect } from 'react';
import { apiRoutes } from "../../plugins/apiRoutes.js";
import { get } from "./../../plugins/request.js";
import {getText, lang} from "../../lang/lang.js";

export default function SelectLanguage({setLang, value = 0, disabled = false, exceptId = 0, extraEmptyField = false}) {
    const [languages, setLanguages] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    function handleLanguage() {
        if(exceptId === 0) {
            get(apiRoutes.languages, null, { withCredentials: true })
                .then(data => {
                    if (Array.isArray(data)) {
                        setLanguages(data);
                    }
                    else if (data?.data && Array.isArray(data.data)) {
                        setLanguages(data.data);
                    }
                    else if (data?.languages && Array.isArray(data.languages)) {
                        setLanguages(data.languages);
                    }
                    setLoading(false);
                })
                .catch(err => {
                    setError(err.message);
                    setLoading(false);
                });
        }
        else {
            get(apiRoutes.exceptLanguage + '/' + exceptId, null, { withCredentials: true })
                .then(data => {
                    if (Array.isArray(data)) {
                        setLanguages(data);
                    }
                    else if (data?.data && Array.isArray(data.data)) {
                        setLanguages(data.data);
                    }
                    else if (data?.languages && Array.isArray(data.languages)) {
                        setLanguages(data.languages);
                    }
                    setLoading(false);
                })
                .catch(err => {
                    setError(err.message);
                    setLoading(false);
                });
        }
    }
    useEffect(() => {
        handleLanguage()
    }, [exceptId]);

    if (loading) {
        return (
            <select className="w-full px-4 py-2.5 rounded-xl border border-slate-200 bg-slate-50 text-slate-400 cursor-not-allowed" disabled>
                <option>{getText(lang.selectedLanguage.loading)}</option>
            </select>
        );
    }

    if (error) {
        return (
            <select className="w-full px-4 py-2.5 rounded-xl border border-rose-200 bg-rose-50 text-rose-500 cursor-not-allowed" disabled>
                <option>{getText(lang.selectedLanguage.errorLoading)}</option>
            </select>
        );
    }

    return (
        <select
            className="w-full px-4 py-2.5 rounded-xl border border-slate-200 bg-white hover:border-slate-300 focus:border-indigo-400 focus:ring-2 focus:ring-indigo-400/20 outline-none transition-all duration-200 text-slate-700 cursor-pointer appearance-none"
            value={value}
            onChange={(e) => setLang(e.target.value)}
            disabled={disabled}
        >
            {extraEmptyField ? (<option value={0}>{getText(lang.selectedLanguage.chooseLanguage)}</option>) : ''}
            {languages.map(lang => (
                <option key={lang.id} value={lang.id}>
                    {lang.name}
                </option>
            ))}
        </select>
    );
}