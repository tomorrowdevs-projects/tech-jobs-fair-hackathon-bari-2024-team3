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
import "./style/style-new-game.css";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { useEffect, useState } from "react";

export default function Home() {
  const [username, setUsername] = useState("");
  const [room, setRoom] = useState("");

  const WS_URL = "ws://localhost:3333/ws";
  const { sendMessage, lastMessage, readyState } = useWebSocket(WS_URL, {});
  const createQuiz = () => {
    if (readyState === ReadyState.OPEN) {
      sendMessage(`createQuiz Quiz ${username}`);
      sendMessage(`setUsername ${username}`);
    }
  };

  useEffect(() => {
    if (lastMessage?.data) {
      const data = lastMessage.data.split(",");
      if (data) {
        const id = data[0].split(":");
        if (id) {
          setRoom(id[1])
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
            placeholder="Username"
            className="input-room-code"
            onChange={(e) => setUsername(e.target.value)}
          />
        </InputGroup>
        <Link
          href={`/new-game/${room?.trimStart() || "33594879-2da1-4fb2-854e-5d6e1f1eb49f"}`}
          onClick={createQuiz}
          style={{ marginBottom: 32 }}
        >
          <Button className="button-one">Enter</Button>
        </Link>
      </Flex>
    </Box>
  );
}
