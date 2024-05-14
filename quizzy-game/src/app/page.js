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
        marginTop="110px" 
      />

      <Flex alignItems="center" flexDirection="column" marginTop="100px">
        <InputGroup width="320px" height="48px">
          <InputLeftElement pointerEvents='none'>
          </InputLeftElement>
          <Input type='rom' placeholder='Room Code' />
        </InputGroup>

        <Button
          size="lg"
          height="48px"
          width="320px"
          border="2px"
          borderColor="green.500"
          marginTop="20px"
          bg="#ECC94B"
        >
          Enter
        </Button>
      </Flex>
    </Box>
  );
}
