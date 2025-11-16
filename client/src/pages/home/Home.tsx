import { Box, Heading, Text } from "@chakra-ui/react";

export default function Home() {
  return (
    <Box textAlign="center" py={20}>
      <Heading>Welcome to Dr. Williams’ Website</Heading>
      <Text mt={4}>Book consultations, read health articles, and access medical courses.</Text>
    </Box>
  );
}
