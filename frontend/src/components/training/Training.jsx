import {useEffect, useState} from "react";
import {studyStatuses} from "../../plugins/studyStatus.js";
import EndTraining from "./EndTraining.jsx";
import WaitingTraining from "./WaitingTraining.jsx";
import Card from "./Card.jsx";
import Loading from "../layouts/Loading.jsx";
export default function Training() {
    const [isTraining, setTraining] = useState(studyStatuses.none)
    return (
        <main
            className="min-h-screen bg-gradient-to-br from-indigo-50 via-purple-50 to-pink-50 flex flex-col items-center justify-center p-6">
            {isTraining === studyStatuses.learning ? (
                <WaitingTraining />
            ) : isTraining === studyStatuses.learned ? (
                <EndTraining/>
            ) : (
                <Card
                   setTraining={setTraining}
                />
            )}
        </main>
    );
}