"use client";
import {
  Box,
  Image,
  Button,
  InputGroup,
  Input,
  InputLeftElement,
  Flex,
} from "@chakra-ui/react";
import Link from "next/link";
import "./page-module.css";
import { useEffect, useState } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";

export default function Home() {
  const [username, setUsername] = useState("");
  const [room, setRoom] = useState("");

  const WS_URL = `ws://localhost:3333/ws/`;
  const { sendMessage, lastMessage, readyState } = useWebSocket(
    WS_URL,
    {},
  );

  const enterRoom = () => {
    if (readyState === ReadyState.OPEN) {
      sendMessage(`setUsername ${username}`)
      sendMessage(`joinQuiz ${room}`);
    }
  };

  useEffect(() => {
    const arr = lastMessage?.data.split(":");
    if (arr) {
      if (arr[0].includes("UserID")) {
        const obj = { UserID: arr[1].trim() };
        if (sessionStorage.getItem("UserID") === null) {
          sessionStorage.setItem("UserID", obj[arr[0]]);
        }
      }
    }
  }, [lastMessage]);
  return (
    <Box className="box">
      <Image
        src="/image-logo.svg"
        alt="Descrizione dell'immagine"
        className="image"
      />
      <Flex alignItems="center" flexDirection="column" marginTop="190px">
        <InputGroup>
          <InputLeftElement pointerEvents="none"></InputLeftElement>
          <Input
            type="rom"
            placeholder="Room Code"
            className="input-room-code"
            onChange={(e) => setRoom(e.target.value)}
          />
        </InputGroup>
        <InputGroup style={{ marginTop: 20 }}>
          <InputLeftElement pointerEvents="none"></InputLeftElement>
          <Input
            type="rom"
            placeholder="Username"
            className="input-room-code"
            onChange={(e) => setUsername(e.target.value)}
          />
        </InputGroup>
        <Link href={`/new-game/${room}`} onClick={enterRoom}>
          <Button className="button-one">Enter</Button>
        </Link>
        <Link href="/new-game">
          <Button className="button-two">+</Button>
        </Link>
      </Flex>
    </Box>
  );
}
