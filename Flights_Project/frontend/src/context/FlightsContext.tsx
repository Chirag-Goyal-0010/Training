import React, { createContext, useState, useContext, useCallback, ReactNode } from 'react';
import { flightAPI, IFlight, ISearchCriteria } from '../api';

// 1. Define Flight Interface
export interface IFlight {
  id: number;
  flightCode: string;
  origin: string;
  destination: string;
  departureTime: string; // ISO string format for date-time
  arrivalTime: string;   // ISO string format for date-time
  capacity: number;
  price: number;
}

// 2. Define Search Criteria Interface
interface ISearchCriteria {
  origin?: string;
  destination?: string;
  date?: string; // Date string (e.g., 'YYYY-MM-DD')
}

// 3. Define FlightsContext Interface
interface IFlightsContext {
  flights: IFlight[];
  loadingFlights: boolean;
  flightsError: string | null;
  fetchFlights: (searchCriteria?: ISearchCriteria) => Promise<void>;
  createFlight: (flightData: Omit<IFlight, 'id'>) => Promise<{ success: boolean; data?: IFlight; error?: string }>;
  updateFlight: (id: number, flightData: Partial<IFlight>) => Promise<{ success: boolean; data?: IFlight; error?: string }>;
  deleteFlight: (id: number) => Promise<{ success: boolean; error?: string }>;
}

interface FlightsProviderProps {
  children: ReactNode;
}

// Initialize Context with a default null value or a mock object matching IFlightsContext
const FlightsContext = createContext<IFlightsContext | undefined>(undefined);

export const FlightsProvider = ({ children }: FlightsProviderProps) => {
  const [flights, setFlights] = useState<IFlight[]>([]);
  const [loadingFlights, setLoadingFlights] = useState<boolean>(false);
  const [flightsError, setFlightsError] = useState<string | null>(null);

  const fetchFlights = useCallback(async (searchCriteria: ISearchCriteria = {}) => {
    setLoadingFlights(true);
    setFlightsError(null);
    try {
      const response = await flightAPI.getFlights(searchCriteria);
      setFlights(response.data);
    } catch (error: any) {
      setFlightsError(error.message || 'Failed to fetch flights');
    } finally {
      setLoadingFlights(false);
    }
  }, []);

  const createFlight = useCallback(async (flightData: Omit<IFlight, 'id'>) => {
    setLoadingFlights(true);
    setFlightsError(null);
    try {
      const response = await flightAPI.createFlight(flightData);
      setFlights((prevFlights: IFlight[]) => [...prevFlights, response.data]);
      return { success: true, data: response.data };
    } catch (error: any) {
      setFlightsError(error.message || 'Failed to create flight');
      return { success: false, error: error.message };
    } finally {
      setLoadingFlights(false);
    }
  }, []);

  const updateFlight = useCallback(async (id: number, flightData: Partial<IFlight>) => {
    setLoadingFlights(true);
    setFlightsError(null);
    try {
      const response = await flightAPI.updateFlight(id, flightData);
      setFlights((prevFlights: IFlight[]) =>
        prevFlights.map((flight: IFlight) => (flight.id === id ? response.data : flight))
      );
      return { success: true, data: response.data };
    } catch (error: any) {
      setFlightsError(error.message || 'Failed to update flight');
      return { success: false, error: error.message };
    } finally {
      setLoadingFlights(false);
    }
  }, []);

  const deleteFlight = useCallback(async (id: number) => {
    setLoadingFlights(true);
    setFlightsError(null);
    try {
      await flightAPI.deleteFlight(id);
      setFlights((prevFlights: IFlight[]) => prevFlights.filter((flight: IFlight) => flight.id !== id));
      return { success: true };
    } catch (error: any) {
      setFlightsError(error.message || 'Failed to delete flight');
      return { success: false, error: error.message };
    } finally {
      setLoadingFlights(false);
    }
  }, []);

  const value: IFlightsContext = {
    flights,
    loadingFlights,
    flightsError,
    fetchFlights,
    createFlight,
    updateFlight,
    deleteFlight,
  };

  return (
    <FlightsContext.Provider value={value}>
      {children}
    </FlightsContext.Provider>
  );
};

export const useFlights = () => {
  const context = useContext(FlightsContext);
  if (!context) {
    throw new Error('useFlights must be used within a FlightsProvider');
  }
  return context as IFlightsContext; // Assert type for consumers
};

export default FlightsContext; 