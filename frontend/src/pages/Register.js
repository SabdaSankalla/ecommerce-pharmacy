import React, { useState } from 'react';
import { Box, Button, FormControl, FormLabel, Input, Heading } from '@chakra-ui/react';
import axios from 'axios';

const Register = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleRegister = async () => {
    try {
      await axios.post('/register', { username, password });
      alert('User  registered successfully');
    } catch (error) {
      alert('Error registering user');
    }
  };

  return (
    <Box p={5}>
      <Heading as="h2" size="lg" color="green.500">Register</Heading>
      <FormControl mt={4}>
        <FormLabel>Username</FormLabel>
        <Input value={username} onChange={(e) => setUsername(e.target.value)} />
        <FormLabel mt={4}>Password</FormLabel>
        <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        <Button mt={4} colorScheme="green" onClick={handleRegister}>Register</Button>
      </FormControl>
    </Box>
  );
};

export default Register;