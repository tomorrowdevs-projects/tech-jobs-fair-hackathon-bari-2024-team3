'user client'
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';
import Link from 'next/link';
import  './style/style-new-game-id.css';
const GamePage = () => {
    const idGame = 123; // Valore statico da rendere dinamico
    return (
        <div>
            {/* <h1>Game ID: {idGame}</h1> */}
            <Box className="box">
                <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image"/>
                <Flex alignItems="center" flexDirection="column" marginTop="100px">
                    <InputGroup>
                    <InputLeftElement pointerEvents='none'></InputLeftElement>
                    <Input type='rom' placeholder='1312lamwlkdaw' className='input-room-code'/>
                    </InputGroup>
                    <Button className='button-one'>Answer 1</Button>
                    <Button className='button-one'>Answer 2</Button>
                    <Button className='button-one'>Answer 3</Button>
                    <Button className='button-one'>Answer 4</Button>
                </Flex>
            </Box>
        </div>
    );
};

export default GamePage;

