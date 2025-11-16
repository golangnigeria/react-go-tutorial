import { Container, Stack } from "@chakra-ui/react";
import TodoForm from "../components/TodoForm";
import TodoList from "../components/TodoList";


export const Task = () => {
  return (
    <Stack h='100vh'>
			<Container>
				<TodoForm />
				<TodoList />
			</Container>
		</Stack>
  )
}
