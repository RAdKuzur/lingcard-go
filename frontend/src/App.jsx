import {useEffect, useState} from 'react'
import './App.css'
import Navbar from "./components/layouts/Navbar.jsx";
import Footer from "./components/layouts/Footer.jsx";
import Login from "./components/auth/Login.jsx";
import {Route, Routes, Navigate} from "react-router-dom";
import News from "./components/home/News.jsx";
import {innerRoutes} from "./plugins/routes.js";
import Training from "./components/training/Training.jsx";
import Dictionary from "./components/dictionary/Dictionary.jsx";
import Progress from "./components/progress/Progress.jsx";
import Profile from "./components/profile/Profile.jsx";
import About from "./components/about/About.jsx";
import ProtectedRoute from "./components/layouts/ProtectedRoute.jsx";
import Register from "./components/auth/Register.jsx";
import UnprotectedRoute from "./components/layouts/UnprotectedRoute.jsx";
import Article from "./components/article/Article.jsx";
import Vote from "./components/vote/Vote.jsx";
import VotePage from "./components/vote/VotePage.jsx";
import CreateArticle from "./components/article/CreateArticle.jsx";
import Support from "./components/support/Support.jsx";
import Home from "./components/home/Home.jsx";
import Knowledge from "./components/knowledge/Knowledge.jsx";
function App() {
    return (
        <div className="flex flex-col min-h-screen">
            <Navbar/>
            <Routes>
                <Route path={innerRoutes.home} element={
                    <Home/>
                }/>
                <Route path={innerRoutes.register} element={
                    <UnprotectedRoute>
                        <Register/>
                    </UnprotectedRoute>
                }/>
                <Route path={innerRoutes.login} element={
                    <UnprotectedRoute>
                        <Login/>
                    </UnprotectedRoute>
                }/>
                <Route path={innerRoutes.all} element={<Navigate to={innerRoutes.home} replace />} />
                <Route path={innerRoutes.news} element={
                    <News/>
                }/>
                <Route path={innerRoutes.articlePath} element={
                    <ProtectedRoute>
                        <Article/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.article} element={
                    <ProtectedRoute>
                        <CreateArticle/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.training} element={
                    <ProtectedRoute>
                        <Training/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.progress} element={
                    <ProtectedRoute>
                        <Progress/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.dictionary} element={
                    <ProtectedRoute>
                        <Dictionary/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.knowledge} element={
                    <Knowledge/>
                }/>
                <Route path={innerRoutes.profile} element={
                    <ProtectedRoute>
                        <Profile/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.votes} element={
                    <ProtectedRoute>
                        <Vote/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.votesPath} element={
                    <ProtectedRoute>
                        <VotePage/>
                    </ProtectedRoute>
                }/>
                <Route path={innerRoutes.about} element={
                    <About/>
                }/>
                <Route path={innerRoutes.support} element={
                    <Support/>
                }/>
            </Routes>
            <Footer/>
        </div>
    )
}

export default App