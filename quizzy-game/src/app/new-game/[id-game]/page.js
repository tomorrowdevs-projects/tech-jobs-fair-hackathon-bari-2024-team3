'use client';
import React, { useState, useEffect } from 'react';
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';
import './style/style-new-game-id.css';

const GamePage = () => {
    const [questions, setQuestions] = useState([]);
    const [loading, setLoading] = useState(true);
    const [currentIndex, setCurrentIndex] = useState(0);

    useEffect(() => {
        const fetchQuestions = async () => {
            try {
                const response = await fetch('https://opentdb.com/api.php?amount=10');
                const data = await response.json();
                if (data.results) {
                    setQuestions(data.results);
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
                        <InputLeftElement pointerEvents="none"></InputLeftElement>
                        <Input
                            type="text"
                            placeholder="Inserisci il codice della stanza"
                            className="input-room-code"
                            value={currentQuestion ? currentQuestion.question : ''}
                            readOnly
                            
                            
                        />
                    </InputGroup>
                    {currentQuestion && (
                        <Box className="question-box">
                            <Flex direction="column">
                                {[...currentQuestion.incorrect_answers, currentQuestion.correct_answer].map((answer, idx) => (
                                    <Button key={idx} className="button-one">{answer}</Button>
                                ))}
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
