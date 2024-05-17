'use client';
import React, { useState, useEffect } from 'react';
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex, Text } from '@chakra-ui/react';
import './style/style-new-game-id.css';
import Link from 'next/link';

const GamePage = () => {
    const [questions, setQuestions] = useState([]);
    const [loading, setLoading] = useState(true);
    const [currentIndex, setCurrentIndex] = useState(0);
    const [clickedButtonIndex, setClickedButtonIndex] = useState(null);
    const [answered, setAnswered] = useState(false);
    const [timer, setTimer] = useState(10); // Imposta il timer a 10 secondi
    const [correctAnswersCount, setCorrectAnswersCount] = useState(0);
    const [quizCompleted, setQuizCompleted] = useState(false);

    const handleButtonClick = (idx) => {
        if (!answered) {
            setClickedButtonIndex(idx);
            setAnswered(true);

            // Verifica se la risposta è corretta
            const isCorrect = currentQuestion.correct_answer === currentQuestion.answers[idx];
            if (isCorrect) {
                setCorrectAnswersCount((prevCount) => prevCount += 1);
            }

            setTimeout(handleNextQuestion, 1000); // Aspetta 1 secondo prima di passare alla prossima domanda
        }
    };

    useEffect(() => {
        const fetchQuestions = async () => {
            try {
                const response = await fetch('https://opentdb.com/api.php?amount=10&category=18&difficulty=easy&type=multiple');
                const data = await response.json();
                if (data.results) {
                    const shuffledQuestions = data.results.map(question => ({
                        ...question,
                        answers: shuffleAnswers([...question.incorrect_answers, question.correct_answer])
                    }));
                    setQuestions(shuffledQuestions);
                } else {
                    console.error('Invalid response structure:', data);
                }
                setLoading(false);
            } catch (error) {
                console.error('Errore nel recupero delle domande:', error);
                setLoading(false);
            }
        };

        fetchQuestions();
    }, []);

    useEffect(() => {
        if (answered || timer <= 0) {
            if (timer <= 0) handleNextQuestion();
            return;
        }

        const interval = setInterval(() => {
            setTimer((prevTimer) => prevTimer - 1);
        }, 1000);

        return () => clearInterval(interval);
    }, [answered, timer]);

    const handleNextQuestion = () => {
        if (currentIndex + 1 < questions.length) {
            setCurrentIndex((prevIndex) => (prevIndex + 1));
            setClickedButtonIndex(null);
            setAnswered(false);
            setTimer(10); // Reimposta il timer a 10 secondi per la prossima domanda
        } else {
            setQuizCompleted(true);
        }
    };

    // Funzione per mischiare le risposte in modo casuale
    const shuffleAnswers = (answers) => {
        const shuffledAnswers = [...answers];
        for (let i = shuffledAnswers.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [shuffledAnswers[i], shuffledAnswers[j]] = [shuffledAnswers[j], shuffledAnswers[i]];
        }
        return shuffledAnswers;
    };

    if (loading) {
        return <div>Caricamento...</div>;
    }

    const currentQuestion = questions[currentIndex];

    if (quizCompleted) {
        return (
            <Box className="box">
                <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image" />
                <Flex alignItems="center" flexDirection="column" marginTop="80px">
                    <Text fontSize="2xl" marginBottom="20px">Quiz completato!</Text>
                    <Text fontSize="2xl" marginBottom="20px">Grazie per aver giocato!</Text>
                    <Text fontSize="lg" marginBottom="10px">Risposte corrette: {correctAnswersCount}</Text>
                    <Text fontSize="lg">Risposte sbagliate: {questions.length - correctAnswersCount}</Text>
                </Flex>
            </Box>
        );
    }

    return (
        <div>
            <Box className="box">
                <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image" />
                <Flex alignItems="center" flexDirection="column" marginTop="20px">
                    <InputGroup marginBottom="20px">
                        <Input
                            className="input-timer"
                            value={timer}
                            readOnly
                            style={{ height: '50px', width: '50px', textAlign: 'center', fontSize: '2rem', marginLeft:'125px' }}
                        />
                    </InputGroup>
                    <InputGroup>
                        <InputLeftElement width="120px"></InputLeftElement>
                        <Input
                            className="input-room-code"
                            value={currentQuestion ? currentQuestion.question : ''}
                            readOnly
                        />
                    </InputGroup>
                    
                    {currentQuestion && (
                        <Box className="question-box">
                            <Flex direction="column">
                                {currentQuestion.answers.map((answer, idx) => {
                                    let buttonClass = '';
                                    if (clickedButtonIndex !== null) {
                                        if (idx === clickedButtonIndex) {
                                            buttonClass = currentQuestion.incorrect_answers.includes(answer) ? 'red-button' : 'green-button';
                                        }
                                    }
                                    return (
                                        <a href="#" key={idx} onClick={(e) => { e.preventDefault(); handleButtonClick(idx); }}>
                                            <button 
                                                className={`button-one ${buttonClass}`} 
                                            >
                                                {answer}
                                            </button>
                                        </a>
                                    );
                                })}
                            </Flex>
                        </Box>
                    )}
                </Flex>
            </Box>
        </div>
    );
};

export default GamePage;
