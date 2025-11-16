import {
  Box,
  Flex,
  Button,
  useColorModeValue,
  useColorMode,
  Text,
  HStack,
  IconButton,
  VStack,
  Drawer,
  DrawerOverlay,
  DrawerContent,
  DrawerHeader,
  DrawerBody,
  DrawerCloseButton,
  useDisclosure,
  Link as ChakraLink,
  Container,
} from "@chakra-ui/react";

import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";
import { FaStethoscope } from "react-icons/fa";
import { HiMenu } from "react-icons/hi";
import { Link } from "react-router-dom";

export default function Navbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const menuItems = [
    { label: "Blog", href: "/blog" },
    { label: "Courses", href: "/courses" },
    { label: "Consultations", href: "/consultations" },
    { label: "Login", href: "/auth/login" },
    { label: "Register", href: "/auth/register" },
  ];

  return (
    <Container maxW={"1100px"}>
      <Box
        bg={useColorModeValue("white", "gray.800")}
        px={4}
        py={2}
        my={4}
        borderRadius={"10"}
        boxShadow="md"
      >
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
          
          {/* BRAND */}
          <HStack spacing={3}>
            <FaStethoscope size={32} color="#2B6CB0" />
            <Flex direction="column">
              <Text fontSize="lg" fontWeight="700">
                Dr. Williams
              </Text>
              <Text fontSize="sm" color={"gray.500"}>
                Personal Health Website
              </Text>
            </Flex>
          </HStack>

          {/* RIGHT */}
          <Flex alignItems={"center"} gap={5}>

            {/* DESKTOP MENU */}
            <HStack spacing={6} display={{ base: "none", md: "flex" }}>
              {menuItems.map((item) => (
                <ChakraLink
                  as={Link}
                  key={item.href}
                  to={item.href}
                  fontWeight={500}
                  _hover={{ color: "blue.500" }}
                >
                  {item.label}
                </ChakraLink>
              ))}
            </HStack>

            {/* DARK MODE */}
            <Button variant="outline" size="sm" onClick={toggleColorMode}>
              {colorMode === "light" ? <IoMoon /> : <LuSun />}
            </Button>

            {/* MOBILE MENU ICON */}
            <IconButton
              display={{ base: "flex", md: "none" }}
              aria-label="Open Menu"
              icon={<HiMenu />}
              onClick={onOpen}
            />
          </Flex>
        </Flex>
      </Box>

      {/* MOBILE DRAWER */}
      <Drawer placement="right" onClose={onClose} isOpen={isOpen}>
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerHeader>Menu</DrawerHeader>

          <DrawerBody>
            <VStack align="start" spacing={5}>
              {menuItems.map((item) => (
                <ChakraLink
                  as={Link}
                  key={item.href}
                  to={item.href}
                  fontSize="lg"
                  onClick={onClose}
                >
                  {item.label}
                </ChakraLink>
              ))}
            </VStack>
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </Container>
  );
}
