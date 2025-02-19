import React, { useEffect, useState } from 'react';
import { Box, Heading, Text } from '@chakra-ui/react';
import { useParams } from 'react-router-dom';
import axios from 'axios';

const ItemDetails = () => {
  const { id } = useParams();
  const [item, setItem] = useState(null);

  useEffect(() => {
    const fetchItem = async () => {
      const response = await axios.get(`/items/${id}`);
      setItem(response.data);
    };
    fetchItem();
  }, [id]);

  if (!item) return <Text>Loading...</Text>;

  return (
    <Box p={5}>
      <Heading as="h2" size="lg" color="green.500">{item.name}</Heading>
      <Text mt={4}>{item.description}</Text>
      <Text mt={2}>Price: ${item.price}</Text>
    </Box>
  );
};

export default ItemDetails;