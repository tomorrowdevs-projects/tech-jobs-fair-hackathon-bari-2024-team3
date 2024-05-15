'user client'
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';
import Link from 'next/link';

export default function Home() {
  
  return (
    <Box className="box">

      <Image src="/image-logo.svg" alt="Descrizione dell'immagine" className="image"/>
      
      <Flex alignItems="center" flexDirection="column" marginTop="190px">
        <InputGroup>
          <InputLeftElement pointerEvents='none'></InputLeftElement>
          <Input type='rom' placeholder='Room Code' className='input-room-code'/>
        </InputGroup>
        <Link href="#">
          <Button
            as="a"
            size="lg"
            height="48px"
            width="320px"
            border="2px"
            borderColor="green.500"
            marginTop="20px"
            bg="#ECC94B"
            style={{ pointerEvents: 'none' }}
            color="black"
          >
            Enter
          </Button>
          </Link>
          <Link href="#">
          <Button
            as="a"
            size="lg"
            height="40px"
            width="40px"
            border="2px"
            borderColor="green.500"
            marginTop="40px"
            bg="#ECC94B"
            marginLeft="280px"
            fontSize="24px"
            lineHeight="40px" 
            style={{ pointerEvents: 'none' }}
            color="black"
          > 
            +
          </Button>
          </Link>
      </Flex>
    </Box>
  );
}
