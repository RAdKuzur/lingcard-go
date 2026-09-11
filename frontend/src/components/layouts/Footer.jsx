import {getText, lang} from "../../lang/lang.js";
import {innerRoutes} from "../../plugins/routes.js";

export default function Footer() {
    return (
        <footer className="bg-white border-t border-slate-200">
            <div className="mx-auto px-4 h-16 flex items-center justify-between text-sm text-slate-500">
                <span>© 2026 LingCard</span>
                <div className="flex items-center gap-6">
                    <a href={innerRoutes.about} className="hover:text-slate-700 transition-colors">
                        {getText(lang.footer.about)}
                    </a>
                    <a href={innerRoutes.support} className="hover:text-slate-700 transition-colors">
                        {getText(lang.footer.support)}
                    </a>
                </div>
            </div>
        </footer>
    );
}