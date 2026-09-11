import {getText, lang} from "../../lang/lang.js";
import ButtonLink from "../layouts/ButtonLink.jsx";
import {donations} from "../../plugins/donation.js";

export default function Support() {
    return (
        <main className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 p-6">
            <div className="max-w-5xl mx-auto space-y-8">
                <div
                    className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-8 transition-all hover:shadow-xl hover:shadow-slate-300/50">
                    <div className="flex items-center gap-3 mb-4">
                        <h1 className="text-2xl font-bold text-slate-800">{getText(lang.support.label)}</h1>
                    </div>
                    <p className="text-slate-600 leading-relaxed text-lg">
                        {getText(lang.support.content)} <br/><br/>
                    </p>
                    <div className={'flex gap-3'}>
                        <ButtonLink link={donations.DA.link} color={'white'} label={donations.DA.label} icon={donations.DA.icon}></ButtonLink>
                        <ButtonLink link={donations.BOOSTY.link} color={'white'} label={donations.BOOSTY.label} icon={donations.BOOSTY.icon}></ButtonLink>
                    </div>
                </div>
                <div
                    className="bg-white rounded-2xl shadow-lg shadow-slate-200/50 p-8 transition-all hover:shadow-xl hover:shadow-slate-300/50">
                    <div className="flex items-center justify-center gap-3 mb-4">
                        <h1 className="text-2xl font-bold text-slate-800 text-center">
                            {getText(lang.support.heroes)}
                        </h1>
                    </div>
                    <p className="text-slate-600 leading-relaxed text-lg mt-12 mb-12">
                        <span className="block">1. RKuzur - LingCard's creator</span>
                        <span className="block">2. Rin - UI/UX, designer</span>
                    </p>
                </div>
            </div>
        </main>
    );
}