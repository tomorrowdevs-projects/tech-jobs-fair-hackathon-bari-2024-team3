'use client'
import { Box, Image, Heading, Avatar, Button, Link } from '@chakra-ui/react'; 
import './page-modules.css'; 


export default function Winners({ winners }) { 
    return ( 
        <Box className="box"> 
        <Image src="/image-logo.svg" alt="Quizzy game logo" className="image" /> 
            {winners.length > 0 && ( 
            <div>
                <Heading>WINNER!</Heading> 
                <h3>{winners[0]}</h3> 
                <Avatar className="winnerOne" /> 
            </div> )} 
            <div className="otherWinners"> 
            {winners.length > 1 && ( 
            <>
            <Heading>2nd</Heading> 
            <h3>{topThreeWinners[1]}</h3> 
            <Avatar className="winnerTwo" /> 
            </> )} 
            {winners.length > 2 && ( 
            <> 
            <Heading>3rd</Heading> 
            <h3>{winners[2]}</h3> 
            <Avatar className="winnerThree" /> 
            </> )} 
            </div> 
            <Link href='#'> 
            <Button className="button">MAIN MENU</Button>
        </Link> 
    </Box> 
    ); 
}
