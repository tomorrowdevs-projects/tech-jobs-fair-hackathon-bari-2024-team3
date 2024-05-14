
import { Box, Image, Button, InputGroup, Input, InputLeftElement, Flex } from '@chakra-ui/react';


export default function Home() {
  
  return (
    <Box
      display="flex"
      flexDirection="column"
      justifyContent="flex-start"
      alignItems="center"
      height="100vh"
    >
      <Image
        src="/image-logo.svg"
        alt="Descrizione dell'immagine"
        width="256px"
        height="256px"
        marginTop="180px" 

      />
      
      <Flex alignItems="center" flexDirection="column" marginTop="190px">
        <InputGroup>
          <InputLeftElement pointerEvents='none'>
          </InputLeftElement>
          <Input type='rom' placeholder='Room Code' width="315px" height="48px" border="2px"/>
        </InputGroup>

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
      </Flex>
    </Box>
  );
}
