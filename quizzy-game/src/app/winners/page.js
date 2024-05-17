'use client'
import { Box, Image, Heading, Avatar, Button, Link} from '@chakra-ui/react';
import './page-modules.css';


export default function Winners({ winners = ['Player 1', 'Player 2', 'Player 3'] }){
    return(
        <Box className="box">
            <Image src="/image-logo.svg" alt="Quizzy game logo" className="image" />
                <div>
                    <Heading>WINNER!</Heading>
                    <h3>{winners[0]}</h3>
                    <Avatar className="winnerOne" />
                </div>
            <div className="otherWinners">
                        <Heading>2nd</Heading>
                        <h3>{winners[1]}</h3>
                        <Avatar className="winnerTwo" />
                        <Heading>3rd</Heading>
                        <h3>{winners[2]}</h3>
                        <Avatar className="winnerThree" />
            </div>
            <Link href='#'>
                <Button className="button">MAIN MENU</Button>
            </Link>
        </Box>
    )
}