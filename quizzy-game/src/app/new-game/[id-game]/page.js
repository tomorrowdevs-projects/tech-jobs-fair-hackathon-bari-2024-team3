'use client';
import React, { useState, useEffect } from 'react';
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';
import './style/style-new-game-id.css';
import Link from 'next/link';

const GamePage = () => {
    const [questions, setQuestions] = useState([]);
    const [loading, setLoading] = useState(true);
    const [currentIndex, setCurrentIndex] = useState(0);
    const [clickedButtonIndex, setClickedButtonIndex] = useState(null);
    const [answered, setAnswered] = useState(false); 

    const handleButtonClick = (idx) => {
        if (!answered) {
            setClickedButtonIndex(idx);
            setAnswered(true); 
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

    const handleNextQuestion = () => {
        setCurrentIndex((prevIndex) => (prevIndex + 1) % questions.length);
        setClickedButtonIndex(null);
        setAnswered(false); 
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

    return (
        <div>
            <Box className="box">
                <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image" />
                <Flex alignItems="center" flexDirection="column" marginTop="80px">
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
                    <Button className="button-next" onClick={handleNextQuestion}>Prossima domanda</Button>
                </Flex>
            </Box>
        </div>
    );
};

export default GamePage;
