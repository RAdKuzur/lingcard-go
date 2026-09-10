--
-- PostgreSQL database dump
--

-- Dumped from database version 17.10 (Debian 17.10-1.pgdg13+1)
-- Dumped by pg_dump version 18.4

-- Started on 2026-09-10 13:50:57

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 249 (class 1259 OID 17083)
-- Name: available_languages; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.available_languages (
    id bigint NOT NULL,
    base_language_id bigint NOT NULL,
    target_language_id bigint NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.available_languages OWNER TO admin;

--
-- TOC entry 248 (class 1259 OID 17082)
-- Name: available_languages_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.available_languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.available_languages_id_seq OWNER TO admin;

--
-- TOC entry 3659 (class 0 OID 0)
-- Dependencies: 248
-- Name: available_languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.available_languages_id_seq OWNED BY public.available_languages.id;


--
-- TOC entry 223 (class 1259 OID 16908)
-- Name: cache; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.cache (
    key character varying(255) NOT NULL,
    value text NOT NULL,
    expiration bigint NOT NULL
);


ALTER TABLE public.cache OWNER TO admin;

--
-- TOC entry 224 (class 1259 OID 16916)
-- Name: cache_locks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.cache_locks (
    key character varying(255) NOT NULL,
    owner character varying(255) NOT NULL,
    expiration bigint NOT NULL
);


ALTER TABLE public.cache_locks OWNER TO admin;

--
-- TOC entry 261 (class 1259 OID 17184)
-- Name: comments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.comments (
    id bigint NOT NULL,
    post_id bigint NOT NULL,
    user_id bigint NOT NULL,
    text text NOT NULL,
    "time" timestamp(0) without time zone NOT NULL,
    is_fixed boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.comments OWNER TO admin;

--
-- TOC entry 260 (class 1259 OID 17183)
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.comments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.comments_id_seq OWNER TO admin;

--
-- TOC entry 3660 (class 0 OID 0)
-- Dependencies: 260
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- TOC entry 239 (class 1259 OID 17017)
-- Name: courses; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.courses (
    id bigint NOT NULL,
    word_translation_id bigint NOT NULL,
    user_id bigint NOT NULL,
    repeat integer DEFAULT 0 NOT NULL,
    status integer DEFAULT 1 NOT NULL,
    last_time_repeated timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.courses OWNER TO admin;

--
-- TOC entry 238 (class 1259 OID 17016)
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.courses_id_seq OWNER TO admin;

--
-- TOC entry 3661 (class 0 OID 0)
-- Dependencies: 238
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- TOC entry 245 (class 1259 OID 17060)
-- Name: error_logs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.error_logs (
    id bigint NOT NULL,
    trace text NOT NULL,
    message text NOT NULL,
    "time" timestamp(0) without time zone NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.error_logs OWNER TO admin;

--
-- TOC entry 244 (class 1259 OID 17059)
-- Name: error_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.error_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.error_logs_id_seq OWNER TO admin;

--
-- TOC entry 3662 (class 0 OID 0)
-- Dependencies: 244
-- Name: error_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.error_logs_id_seq OWNED BY public.error_logs.id;


--
-- TOC entry 229 (class 1259 OID 16942)
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.failed_jobs (
    id bigint NOT NULL,
    uuid character varying(255) NOT NULL,
    connection character varying(255) NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.failed_jobs OWNER TO admin;

--
-- TOC entry 228 (class 1259 OID 16941)
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.failed_jobs_id_seq OWNER TO admin;

--
-- TOC entry 3663 (class 0 OID 0)
-- Dependencies: 228
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- TOC entry 227 (class 1259 OID 16934)
-- Name: job_batches; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.job_batches (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    total_jobs integer NOT NULL,
    pending_jobs integer NOT NULL,
    failed_jobs integer NOT NULL,
    failed_job_ids text NOT NULL,
    options text,
    cancelled_at integer,
    created_at integer NOT NULL,
    finished_at integer
);


ALTER TABLE public.job_batches OWNER TO admin;

--
-- TOC entry 226 (class 1259 OID 16925)
-- Name: jobs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.jobs (
    id bigint NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    attempts smallint NOT NULL,
    reserved_at integer,
    available_at integer NOT NULL,
    created_at integer NOT NULL
);


ALTER TABLE public.jobs OWNER TO admin;

--
-- TOC entry 225 (class 1259 OID 16924)
-- Name: jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.jobs_id_seq OWNER TO admin;

--
-- TOC entry 3664 (class 0 OID 0)
-- Dependencies: 225
-- Name: jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.jobs_id_seq OWNED BY public.jobs.id;


--
-- TOC entry 233 (class 1259 OID 16968)
-- Name: languages; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.languages (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    code character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    is_active boolean DEFAULT false NOT NULL
);


ALTER TABLE public.languages OWNER TO admin;

--
-- TOC entry 232 (class 1259 OID 16967)
-- Name: languages_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.languages_id_seq OWNER TO admin;

--
-- TOC entry 3665 (class 0 OID 0)
-- Dependencies: 232
-- Name: languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.languages_id_seq OWNED BY public.languages.id;


--
-- TOC entry 218 (class 1259 OID 16386)
-- Name: migrations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.migrations (
    id integer NOT NULL,
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE public.migrations OWNER TO admin;

--
-- TOC entry 217 (class 1259 OID 16385)
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.migrations_id_seq OWNER TO admin;

--
-- TOC entry 3666 (class 0 OID 0)
-- Dependencies: 217
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- TOC entry 221 (class 1259 OID 16892)
-- Name: password_reset_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.password_reset_tokens (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE public.password_reset_tokens OWNER TO admin;

--
-- TOC entry 231 (class 1259 OID 16955)
-- Name: personal_access_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.personal_access_tokens (
    id bigint NOT NULL,
    tokenable_type character varying(255) NOT NULL,
    tokenable_id bigint NOT NULL,
    name text NOT NULL,
    token character varying(64) NOT NULL,
    abilities text,
    last_used_at timestamp(0) without time zone,
    expires_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.personal_access_tokens OWNER TO admin;

--
-- TOC entry 230 (class 1259 OID 16954)
-- Name: personal_access_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.personal_access_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.personal_access_tokens_id_seq OWNER TO admin;

--
-- TOC entry 3667 (class 0 OID 0)
-- Dependencies: 230
-- Name: personal_access_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.personal_access_tokens_id_seq OWNED BY public.personal_access_tokens.id;


--
-- TOC entry 247 (class 1259 OID 17069)
-- Name: posts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    content text NOT NULL,
    date timestamp(0) without time zone NOT NULL,
    language_id bigint NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    user_id bigint NOT NULL,
    address character varying(255),
    status integer DEFAULT 2 NOT NULL,
    views_count integer DEFAULT 0 NOT NULL,
    likes_count integer DEFAULT 0 NOT NULL,
    dislikes_count integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.posts OWNER TO admin;

--
-- TOC entry 246 (class 1259 OID 17068)
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.posts_id_seq OWNER TO admin;

--
-- TOC entry 3668 (class 0 OID 0)
-- Dependencies: 246
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- TOC entry 253 (class 1259 OID 17126)
-- Name: reactions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.reactions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    status integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.reactions OWNER TO admin;

--
-- TOC entry 252 (class 1259 OID 17125)
-- Name: reactions_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.reactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reactions_id_seq OWNER TO admin;

--
-- TOC entry 3669 (class 0 OID 0)
-- Dependencies: 252
-- Name: reactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.reactions_id_seq OWNED BY public.reactions.id;


--
-- TOC entry 222 (class 1259 OID 16899)
-- Name: sessions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sessions (
    id character varying(255) NOT NULL,
    user_id bigint,
    ip_address character varying(45),
    user_agent text,
    payload text NOT NULL,
    last_activity integer NOT NULL
);


ALTER TABLE public.sessions OWNER TO admin;

--
-- TOC entry 251 (class 1259 OID 17108)
-- Name: suggestions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.suggestions (
    id bigint NOT NULL,
    message text,
    date timestamp(0) without time zone,
    user_id bigint NOT NULL,
    status boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.suggestions OWNER TO admin;

--
-- TOC entry 250 (class 1259 OID 17107)
-- Name: suggestions_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.suggestions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.suggestions_id_seq OWNER TO admin;

--
-- TOC entry 3670 (class 0 OID 0)
-- Dependencies: 250
-- Name: suggestions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.suggestions_id_seq OWNED BY public.suggestions.id;


--
-- TOC entry 241 (class 1259 OID 17036)
-- Name: tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tokens (
    id bigint NOT NULL,
    refresh_token text NOT NULL,
    user_id bigint NOT NULL,
    ip_address character varying(255) NOT NULL,
    user_agent character varying(255) NOT NULL,
    is_revoked boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    expires_at timestamp(0) without time zone
);


ALTER TABLE public.tokens OWNER TO admin;

--
-- TOC entry 240 (class 1259 OID 17035)
-- Name: tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tokens_id_seq OWNER TO admin;

--
-- TOC entry 3671 (class 0 OID 0)
-- Dependencies: 240
-- Name: tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tokens_id_seq OWNED BY public.tokens.id;


--
-- TOC entry 220 (class 1259 OID 16882)
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255),
    email_verified_at timestamp(0) without time zone,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    base_language_id bigint NOT NULL,
    target_language_id bigint NOT NULL,
    role integer DEFAULT 2 NOT NULL,
    is_banned boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- TOC entry 219 (class 1259 OID 16881)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO admin;

--
-- TOC entry 3672 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 243 (class 1259 OID 17051)
-- Name: visits; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.visits (
    id bigint NOT NULL,
    path character varying(255) NOT NULL,
    ip character varying(255) NOT NULL,
    user_agent character varying(255) NOT NULL,
    "time" timestamp(0) without time zone NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    country character varying(255),
    code character varying(255),
    city character varying(255)
);


ALTER TABLE public.visits OWNER TO admin;

--
-- TOC entry 242 (class 1259 OID 17050)
-- Name: visits_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.visits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.visits_id_seq OWNER TO admin;

--
-- TOC entry 3673 (class 0 OID 0)
-- Dependencies: 242
-- Name: visits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.visits_id_seq OWNED BY public.visits.id;


--
-- TOC entry 259 (class 1259 OID 17167)
-- Name: voices; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.voices (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    vote_option_id bigint NOT NULL,
    "time" timestamp(0) without time zone NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.voices OWNER TO admin;

--
-- TOC entry 258 (class 1259 OID 17166)
-- Name: voices_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.voices_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.voices_id_seq OWNER TO admin;

--
-- TOC entry 3674 (class 0 OID 0)
-- Dependencies: 258
-- Name: voices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.voices_id_seq OWNED BY public.voices.id;


--
-- TOC entry 257 (class 1259 OID 17153)
-- Name: vote_options; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.vote_options (
    id bigint NOT NULL,
    vote_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    content character varying(255),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.vote_options OWNER TO admin;

--
-- TOC entry 256 (class 1259 OID 17152)
-- Name: vote_options_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.vote_options_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.vote_options_id_seq OWNER TO admin;

--
-- TOC entry 3675 (class 0 OID 0)
-- Dependencies: 256
-- Name: vote_options_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.vote_options_id_seq OWNED BY public.vote_options.id;


--
-- TOC entry 255 (class 1259 OID 17143)
-- Name: votes; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.votes (
    id bigint NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    is_active boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.votes OWNER TO admin;

--
-- TOC entry 254 (class 1259 OID 17142)
-- Name: votes_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.votes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.votes_id_seq OWNER TO admin;

--
-- TOC entry 3676 (class 0 OID 0)
-- Dependencies: 254
-- Name: votes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.votes_id_seq OWNED BY public.votes.id;


--
-- TOC entry 237 (class 1259 OID 17000)
-- Name: word_translations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.word_translations (
    id bigint NOT NULL,
    word_id bigint NOT NULL,
    target_language_id bigint NOT NULL,
    translation character varying(255),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.word_translations OWNER TO admin;

--
-- TOC entry 236 (class 1259 OID 16999)
-- Name: word_translations_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.word_translations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.word_translations_id_seq OWNER TO admin;

--
-- TOC entry 3677 (class 0 OID 0)
-- Dependencies: 236
-- Name: word_translations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.word_translations_id_seq OWNED BY public.word_translations.id;


--
-- TOC entry 235 (class 1259 OID 16988)
-- Name: words; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.words (
    id bigint NOT NULL,
    text character varying(255) NOT NULL,
    language_id bigint NOT NULL,
    level integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    transcription character varying(255)
);


ALTER TABLE public.words OWNER TO admin;

--
-- TOC entry 234 (class 1259 OID 16987)
-- Name: words_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.words_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.words_id_seq OWNER TO admin;

--
-- TOC entry 3678 (class 0 OID 0)
-- Dependencies: 234
-- Name: words_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.words_id_seq OWNED BY public.words.id;


--
-- TOC entry 3413 (class 2604 OID 17086)
-- Name: available_languages id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.available_languages ALTER COLUMN id SET DEFAULT nextval('public.available_languages_id_seq'::regclass);


--
-- TOC entry 3421 (class 2604 OID 17187)
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- TOC entry 3401 (class 2604 OID 17020)
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- TOC entry 3407 (class 2604 OID 17063)
-- Name: error_logs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.error_logs ALTER COLUMN id SET DEFAULT nextval('public.error_logs_id_seq'::regclass);


--
-- TOC entry 3394 (class 2604 OID 16945)
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- TOC entry 3393 (class 2604 OID 16928)
-- Name: jobs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public.jobs_id_seq'::regclass);


--
-- TOC entry 3397 (class 2604 OID 16971)
-- Name: languages id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.languages ALTER COLUMN id SET DEFAULT nextval('public.languages_id_seq'::regclass);


--
-- TOC entry 3389 (class 2604 OID 16389)
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- TOC entry 3396 (class 2604 OID 16958)
-- Name: personal_access_tokens id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.personal_access_tokens ALTER COLUMN id SET DEFAULT nextval('public.personal_access_tokens_id_seq'::regclass);


--
-- TOC entry 3408 (class 2604 OID 17072)
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- TOC entry 3416 (class 2604 OID 17129)
-- Name: reactions id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reactions ALTER COLUMN id SET DEFAULT nextval('public.reactions_id_seq'::regclass);


--
-- TOC entry 3414 (class 2604 OID 17111)
-- Name: suggestions id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suggestions ALTER COLUMN id SET DEFAULT nextval('public.suggestions_id_seq'::regclass);


--
-- TOC entry 3404 (class 2604 OID 17039)
-- Name: tokens id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tokens ALTER COLUMN id SET DEFAULT nextval('public.tokens_id_seq'::regclass);


--
-- TOC entry 3390 (class 2604 OID 16885)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3406 (class 2604 OID 17054)
-- Name: visits id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.visits ALTER COLUMN id SET DEFAULT nextval('public.visits_id_seq'::regclass);


--
-- TOC entry 3420 (class 2604 OID 17170)
-- Name: voices id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.voices ALTER COLUMN id SET DEFAULT nextval('public.voices_id_seq'::regclass);


--
-- TOC entry 3419 (class 2604 OID 17156)
-- Name: vote_options id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.vote_options ALTER COLUMN id SET DEFAULT nextval('public.vote_options_id_seq'::regclass);


--
-- TOC entry 3417 (class 2604 OID 17146)
-- Name: votes id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.votes ALTER COLUMN id SET DEFAULT nextval('public.votes_id_seq'::regclass);


--
-- TOC entry 3400 (class 2604 OID 17003)
-- Name: word_translations id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.word_translations ALTER COLUMN id SET DEFAULT nextval('public.word_translations_id_seq'::regclass);


--
-- TOC entry 3399 (class 2604 OID 16991)
-- Name: words id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.words ALTER COLUMN id SET DEFAULT nextval('public.words_id_seq'::regclass);


--
-- TOC entry 3476 (class 2606 OID 17088)
-- Name: available_languages available_languages_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.available_languages
    ADD CONSTRAINT available_languages_pkey PRIMARY KEY (id);


--
-- TOC entry 3440 (class 2606 OID 16922)
-- Name: cache_locks cache_locks_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cache_locks
    ADD CONSTRAINT cache_locks_pkey PRIMARY KEY (key);


--
-- TOC entry 3437 (class 2606 OID 16914)
-- Name: cache cache_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cache
    ADD CONSTRAINT cache_pkey PRIMARY KEY (key);


--
-- TOC entry 3488 (class 2606 OID 17192)
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- TOC entry 3466 (class 2606 OID 17024)
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- TOC entry 3472 (class 2606 OID 17067)
-- Name: error_logs error_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.error_logs
    ADD CONSTRAINT error_logs_pkey PRIMARY KEY (id);


--
-- TOC entry 3448 (class 2606 OID 16950)
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- TOC entry 3450 (class 2606 OID 16953)
-- Name: failed_jobs failed_jobs_uuid_unique; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_uuid_unique UNIQUE (uuid);


--
-- TOC entry 3445 (class 2606 OID 16940)
-- Name: job_batches job_batches_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.job_batches
    ADD CONSTRAINT job_batches_pkey PRIMARY KEY (id);


--
-- TOC entry 3442 (class 2606 OID 16932)
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);


--
-- TOC entry 3458 (class 2606 OID 16975)
-- Name: languages languages_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


--
-- TOC entry 3424 (class 2606 OID 16391)
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- TOC entry 3430 (class 2606 OID 16898)
-- Name: password_reset_tokens password_reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT password_reset_tokens_pkey PRIMARY KEY (email);


--
-- TOC entry 3453 (class 2606 OID 16962)
-- Name: personal_access_tokens personal_access_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.personal_access_tokens
    ADD CONSTRAINT personal_access_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 3455 (class 2606 OID 16965)
-- Name: personal_access_tokens personal_access_tokens_token_unique; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.personal_access_tokens
    ADD CONSTRAINT personal_access_tokens_token_unique UNIQUE (token);


--
-- TOC entry 3474 (class 2606 OID 17076)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- TOC entry 3480 (class 2606 OID 17131)
-- Name: reactions reactions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reactions
    ADD CONSTRAINT reactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3433 (class 2606 OID 16905)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3478 (class 2606 OID 17116)
-- Name: suggestions suggestions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suggestions
    ADD CONSTRAINT suggestions_pkey PRIMARY KEY (id);


--
-- TOC entry 3468 (class 2606 OID 17044)
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 3426 (class 2606 OID 17206)
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- TOC entry 3428 (class 2606 OID 16889)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3470 (class 2606 OID 17058)
-- Name: visits visits_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.visits
    ADD CONSTRAINT visits_pkey PRIMARY KEY (id);


--
-- TOC entry 3486 (class 2606 OID 17172)
-- Name: voices voices_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.voices
    ADD CONSTRAINT voices_pkey PRIMARY KEY (id);


--
-- TOC entry 3484 (class 2606 OID 17160)
-- Name: vote_options vote_options_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.vote_options
    ADD CONSTRAINT vote_options_pkey PRIMARY KEY (id);


--
-- TOC entry 3482 (class 2606 OID 17151)
-- Name: votes votes_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT votes_pkey PRIMARY KEY (id);


--
-- TOC entry 3463 (class 2606 OID 17005)
-- Name: word_translations word_translations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.word_translations
    ADD CONSTRAINT word_translations_pkey PRIMARY KEY (id);


--
-- TOC entry 3461 (class 2606 OID 16993)
-- Name: words words_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.words
    ADD CONSTRAINT words_pkey PRIMARY KEY (id);


--
-- TOC entry 3435 (class 1259 OID 16915)
-- Name: cache_expiration_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX cache_expiration_index ON public.cache USING btree (expiration);


--
-- TOC entry 3438 (class 1259 OID 16923)
-- Name: cache_locks_expiration_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX cache_locks_expiration_index ON public.cache_locks USING btree (expiration);


--
-- TOC entry 3446 (class 1259 OID 16951)
-- Name: failed_jobs_connection_queue_failed_at_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX failed_jobs_connection_queue_failed_at_index ON public.failed_jobs USING btree (connection, queue, failed_at);


--
-- TOC entry 3443 (class 1259 OID 16933)
-- Name: jobs_queue_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX jobs_queue_index ON public.jobs USING btree (queue);


--
-- TOC entry 3451 (class 1259 OID 16966)
-- Name: personal_access_tokens_expires_at_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX personal_access_tokens_expires_at_index ON public.personal_access_tokens USING btree (expires_at);


--
-- TOC entry 3456 (class 1259 OID 16963)
-- Name: personal_access_tokens_tokenable_type_tokenable_id_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX personal_access_tokens_tokenable_type_tokenable_id_index ON public.personal_access_tokens USING btree (tokenable_type, tokenable_id);


--
-- TOC entry 3431 (class 1259 OID 16907)
-- Name: sessions_last_activity_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX sessions_last_activity_index ON public.sessions USING btree (last_activity);


--
-- TOC entry 3434 (class 1259 OID 16906)
-- Name: sessions_user_id_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX sessions_user_id_index ON public.sessions USING btree (user_id);


--
-- TOC entry 3464 (class 1259 OID 17208)
-- Name: word_translations_target_language_id_word_id_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX word_translations_target_language_id_word_id_index ON public.word_translations USING btree (target_language_id, word_id);


--
-- TOC entry 3459 (class 1259 OID 17207)
-- Name: words_language_id_index; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX words_language_id_index ON public.words USING btree (language_id);


--
-- TOC entry 3499 (class 2606 OID 17089)
-- Name: available_languages available_languages_base_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.available_languages
    ADD CONSTRAINT available_languages_base_language_id_foreign FOREIGN KEY (base_language_id) REFERENCES public.languages(id);


--
-- TOC entry 3500 (class 2606 OID 17094)
-- Name: available_languages available_languages_target_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.available_languages
    ADD CONSTRAINT available_languages_target_language_id_foreign FOREIGN KEY (target_language_id) REFERENCES public.languages(id);


--
-- TOC entry 3507 (class 2606 OID 17193)
-- Name: comments comments_post_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_post_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- TOC entry 3508 (class 2606 OID 17198)
-- Name: comments comments_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3494 (class 2606 OID 17030)
-- Name: courses courses_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3495 (class 2606 OID 17025)
-- Name: courses courses_word_translation_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_word_translation_id_foreign FOREIGN KEY (word_translation_id) REFERENCES public.word_translations(id);


--
-- TOC entry 3497 (class 2606 OID 17077)
-- Name: posts posts_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_language_id_foreign FOREIGN KEY (language_id) REFERENCES public.languages(id);


--
-- TOC entry 3498 (class 2606 OID 17100)
-- Name: posts posts_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3502 (class 2606 OID 17137)
-- Name: reactions reactions_post_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reactions
    ADD CONSTRAINT reactions_post_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- TOC entry 3503 (class 2606 OID 17132)
-- Name: reactions reactions_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reactions
    ADD CONSTRAINT reactions_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3501 (class 2606 OID 17117)
-- Name: suggestions suggestions_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suggestions
    ADD CONSTRAINT suggestions_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3496 (class 2606 OID 17045)
-- Name: tokens tokens_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3489 (class 2606 OID 16976)
-- Name: users users_base_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_base_language_id_foreign FOREIGN KEY (base_language_id) REFERENCES public.languages(id);


--
-- TOC entry 3490 (class 2606 OID 16981)
-- Name: users users_target_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_target_language_id_foreign FOREIGN KEY (target_language_id) REFERENCES public.languages(id);


--
-- TOC entry 3505 (class 2606 OID 17173)
-- Name: voices voices_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.voices
    ADD CONSTRAINT voices_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3506 (class 2606 OID 17178)
-- Name: voices voices_vote_option_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.voices
    ADD CONSTRAINT voices_vote_option_id_foreign FOREIGN KEY (vote_option_id) REFERENCES public.vote_options(id);


--
-- TOC entry 3504 (class 2606 OID 17161)
-- Name: vote_options vote_options_vote_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.vote_options
    ADD CONSTRAINT vote_options_vote_id_foreign FOREIGN KEY (vote_id) REFERENCES public.votes(id);


--
-- TOC entry 3492 (class 2606 OID 17011)
-- Name: word_translations word_translations_target_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.word_translations
    ADD CONSTRAINT word_translations_target_language_id_foreign FOREIGN KEY (target_language_id) REFERENCES public.languages(id);


--
-- TOC entry 3493 (class 2606 OID 17006)
-- Name: word_translations word_translations_word_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.word_translations
    ADD CONSTRAINT word_translations_word_id_foreign FOREIGN KEY (word_id) REFERENCES public.words(id);


--
-- TOC entry 3491 (class 2606 OID 16994)
-- Name: words words_language_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.words
    ADD CONSTRAINT words_language_id_foreign FOREIGN KEY (language_id) REFERENCES public.languages(id);


-- Completed on 2026-09-10 13:50:57

--
-- PostgreSQL database dump complete
--

