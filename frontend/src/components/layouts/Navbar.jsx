import { Link, useLocation } from 'react-router-dom';
import ProfileBar from "./ProfileBar.jsx";
import Logo from "./Logo.jsx";
import {innerRoutes} from "../../plugins/routes.js";
import {useAuth} from "../../plugins/AuthContext.jsx";
import {useEffect, useState, useRef} from "react";
import {getText, lang, languageOptions} from "../../lang/lang.js";
import ArrowDown from "../svg/ArrowDown.jsx";
import Burger from "../svg/Burger.jsx";
import Choose from "../svg/Choose.jsx";

export default function Navbar() {
    const auth = useAuth();
    const location = useLocation();
    const [currentLang, setCurrentLang] = useState('en');
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
    const [isLanguageDropdownOpen, setIsLanguageDropdownOpen] = useState(false);
    const languageDropdownRef = useRef(null);

    const menuOptions = {
        news: {
            link: innerRoutes.news,
            label: getText(lang.navbar.options.news)
        },
        training: {
            link: innerRoutes.training,
            label: getText(lang.navbar.options.training)
        },
        progress: {
            link: innerRoutes.progress,
            label: getText(lang.navbar.options.progress)
        },
        dictionary: {
            link: innerRoutes.dictionary,
            label: getText(lang.navbar.options.dictionary)
        },
        newLanguage: {
            link: innerRoutes.votes,
            label: getText(lang.navbar.options.votes)
        },
        about: {
            link: innerRoutes.about,
            label: getText(lang.navbar.options.about)
        },
        support: {
            link: innerRoutes.support,
            label: getText(lang.navbar.options.support)
        },
        knowledge: {
            link: innerRoutes.knowledge,
            label: getText(lang.navbar.options.knowledge)
        },
        profile: {
            link: innerRoutes.profile,
            label: getText(lang.navbar.options.profile)
        }
    }

    const menuGuestOptions = {
        about: {
            link: innerRoutes.about,
            label: getText(lang.navbar.options.about)
        },
        support: {
            link: innerRoutes.support,
            label: getText(lang.navbar.options.support)
        },
        knowledge: {
            link: innerRoutes.knowledge,
            label: getText(lang.navbar.options.knowledge)
        },
    }

    const isActivePath = (path) => {
        if (path === innerRoutes.news) {
            return location.pathname === path;
        }
        return location.pathname.startsWith(path);
    };

    useEffect(() => {
        const language = localStorage.getItem('lang') ?? 'en';
        setCurrentLang(language);
        localStorage.setItem('lang', language);
    }, []);

    useEffect(() => {
        const handleClickOutside = (e) => {
            if (languageDropdownRef.current && !languageDropdownRef.current.contains(e.target)) {
                setIsLanguageDropdownOpen(false);
            }
        };
        document.addEventListener('click', handleClickOutside);
        return () => document.removeEventListener('click', handleClickOutside);
    }, []);

    useEffect(() => {
        const handleClickOutside = (e) => {
            if (isMobileMenuOpen && !e.target.closest('nav')) {
                setIsMobileMenuOpen(false);
            }
        };
        document.addEventListener('click', handleClickOutside);
        return () => document.removeEventListener('click', handleClickOutside);
    }, [isMobileMenuOpen]);

    useEffect(() => {
        if (isMobileMenuOpen) {
            document.body.style.overflow = 'hidden';
        } else {
            document.body.style.overflow = 'unset';
        }
        return () => {
            document.body.style.overflow = 'unset';
        };
    }, [isMobileMenuOpen]);

    const handleLanguageChange = (value) => {
        setCurrentLang(value);
        localStorage.setItem('lang', value);
        setIsLanguageDropdownOpen(false);
        window.location.reload();
    };

    const selectedLanguage = languageOptions.find(l => l.value === currentLang);

    return (
        <>
            <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-lg border-b border-slate-200/50 shadow-sm">
                <div className="max-w-12xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div className="flex justify-between items-center h-16 md:h-20">
                        <div className="flex-shrink-0">
                            <Logo />
                        </div>
                        <div className="hidden md:flex items-center gap-4 lg:gap-8">
                            {auth.isAuthenticated() ? Object.values(menuOptions).map((item) => {
                                const isActive = isActivePath(item.link);
                                return (
                                    <Link
                                        key={item.label}
                                        to={item.link}
                                        className={`text-sm font-medium transition-colors duration-200 
                                            pb-1 whitespace-nowrap
                                            ${isActive
                                            ? 'text-emerald-600 border-b-2 border-emerald-600'
                                            : 'text-slate-700 hover:text-emerald-600 border-b-2 border-transparent hover:border-emerald-600'
                                        }`}
                                    >
                                        {item.label}
                                    </Link>
                                );
                            }) :
                                Object.values(menuGuestOptions).map((item) => {
                                    const isActive = isActivePath(item.link);
                                    return (
                                        <Link
                                            key={item.label}
                                            to={item.link}
                                            className={`text-sm font-medium transition-colors duration-200 
                                            pb-1 whitespace-nowrap
                                            ${isActive
                                                ? 'text-emerald-600 border-b-2 border-emerald-600'
                                                : 'text-slate-700 hover:text-emerald-600 border-b-2 border-transparent hover:border-emerald-600'
                                            }`}
                                        >
                                            {item.label}
                                        </Link>
                                    );
                                })

                            }
                        </div>

                        <div className="flex items-center gap-2 sm:gap-4">
                            <div className="relative" ref={languageDropdownRef}>
                                <button
                                    onClick={() => setIsLanguageDropdownOpen(!isLanguageDropdownOpen)}
                                    className="flex items-center gap-2 border border-slate-200 rounded-lg
                                             pl-2 sm:pl-2 pr-2 sm:pr-2 py-1 sm:py-1.5
                                             text-xs sm:text-sm font-medium text-slate-700
                                             hover:border-indigo-400 focus:outline-none focus:ring-2
                                             focus:ring-indigo-500/20 focus:border-indigo-500
                                             transition-all duration-200 cursor-pointer
                                             min-w-[60px] sm:min-w-[80px] md:min-w-[120px]
                                             bg-transparent"
                                >
                                    <img
                                        src={selectedLanguage?.flag}
                                        alt={selectedLanguage?.name}
                                        className="w-4 h-4 sm:w-5 sm:h-5 rounded-sm object-cover"
                                    />
                                    <span className="hidden md:inline">{selectedLanguage?.name}</span>
                                    <ArrowDown/>
                                </button>

                                {isLanguageDropdownOpen && (
                                    <div className="absolute right-0 mt-1 min-w-[160px] bg-white
                                                  border border-slate-200 rounded-lg shadow-lg
                                                  py-1 z-50 animate-fadeIn">
                                        {languageOptions.map((lang) => (
                                            <button
                                                key={lang.value}
                                                onClick={() => handleLanguageChange(lang.value)}
                                                className="w-full flex items-center gap-3 px-4 py-2
                                                         text-sm text-slate-700 hover:bg-indigo-50
                                                         transition-colors duration-150"
                                            >
                                                <img
                                                    src={lang.flag}
                                                    alt={lang.name}
                                                    className="w-5 h-5 rounded-sm object-cover"
                                                />
                                                <span>{lang.name}</span>
                                                {lang.value === currentLang && (
                                                    <Choose/>
                                                )}
                                            </button>
                                        ))}
                                    </div>
                                )}
                            </div>
                            {(!auth.isAuthenticated() && !auth.isLoading) ? <ProfileBar/> : <></>}
                            {auth.isAuthenticated() ? (
                                <button
                                    onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                                    className="md:hidden p-2 rounded-lg hover:bg-slate-100 transition-colors duration-200"
                                    aria-label="Toggle menu"
                                >
                                    <Burger isMobileMenuOpen={isMobileMenuOpen}/>
                                </button>
                            ) : (<></>)}
                        </div>
                    </div>
                </div>
            </nav>
            {isMobileMenuOpen && auth.isAuthenticated() && (
                <div className="md:hidden fixed inset-0 z-40 bg-black/20 backdrop-blur-sm animate-fadeIn">
                    <div className="fixed inset-x-0 top-16 bg-white shadow-xl border-b border-slate-200 animate-slideDown">
                        <div className="px-4 py-3 space-y-1 max-h-[calc(100vh-4rem)] overflow-y-auto">
                            {Object.values(menuOptions).map((item, index) => {
                                const isActive = isActivePath(item.link);
                                return (
                                    <Link
                                        key={item.label}
                                        to={item.link}
                                        onClick={() => setIsMobileMenuOpen(false)}
                                        className={`block px-4 py-3 rounded-lg text-base font-medium
                                             transition-all duration-200 border-l-4
                                             ${isActive
                                            ? 'text-indigo-600 bg-indigo-50 border-indigo-500'
                                            : 'text-slate-700 hover:bg-indigo-50 hover:text-indigo-600 border-transparent hover:border-indigo-500'
                                        }`}
                                        style={{
                                            animationDelay: `${index * 50}ms`
                                        }}
                                    >
                                        {item.label}
                                    </Link>
                                );
                            })}
                        </div>
                    </div>
                </div>
            )}
        </>
    );
}