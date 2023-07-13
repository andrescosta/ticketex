import { Address } from './address';

interface Venue {
  name: string;
  address: Address;
  phone: string;
}

interface Ticket {
  capacity: number;
  type: string;
}

export interface Adventure {
  name: string;
  description: string;
  tickets: Array<Ticket>;
  venue: Venue;
}
