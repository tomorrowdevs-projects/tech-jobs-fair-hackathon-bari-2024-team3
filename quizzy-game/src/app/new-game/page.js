
'user client'
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';
import Link from 'next/link';
import  './style/style-new-game.css';



export default function Home() {
  return (
    <Box className="box">
      <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image"/>
      <Flex alignItems="center" flexDirection="column" marginTop="190px">
        <InputGroup>
          <InputLeftElement pointerEvents='none'></InputLeftElement>
          <Input type='rom' placeholder='1312lamwlkdaw' className='input-room-code'/>
        </InputGroup>
        <Link href="#">
          <Button className='button-one'>Enter</Button>
        </Link>
        <Link href="#">
          <Button className='button-two'>+</Button>
        </Link>
      </Flex>
    </Box>
  );
}
