'use client'
import { Box, Image, Table, Thead, Tbody, Tr, Th, Td, TableContainer, Button} from '@chakra-ui/react';
import './page-modules.css';



export default function NewGameWaiting({ players = ['Player 1', 'Player 2'] }) {
    
    return (
     
      <Box className="box">
        <Image src="/image-logo.svg" alt="Quizzy game logo" className="image"/>

        <TableContainer className="tableContainer" >
          <Table variant='simple' className="table">
            <Thead>
              <Tr>
                <Th className="titleTable">PLAYER</Th>
                <Th>AVATAR</Th>
              </Tr>
            </Thead>
            {players.map((players, index) =>(
              <Tbody key={index}>
              <Tr>
                <Td>{players}</Td>
                <Td>pass the value</Td>
              </Tr>
            </Tbody>
              ))}
          </Table>
        </TableContainer>
        <Button
          isLoading
          colorScheme='yellow'
        >
          Waiting for game start...
        </Button>

      </Box>
      
    );
  }