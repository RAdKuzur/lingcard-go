import { innerRoutes } from "../../plugins/routes.js";
import {useRedirect} from "../../hooks/useRedirect.js";

export default function Logo() {
    const {redirect} = useRedirect();

    function goHome() {
        redirect(innerRoutes.home);
    }
    return (
        <div
            className="flex items-center gap-2 cursor-pointer transition-all duration-300 hover:scale-105" onClick={goHome}>
            <img className={'w-20 h-20'} src={'/icons/Logo.svg'} alt={'LingCard'}/>
        </div>
    );
}