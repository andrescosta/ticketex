import { Address } from './address';

export interface User {
  id: string;
  email: string;
  fullname: string;
  phone: string;
  extId: string;
  preferences: Array<Array<Channel>>;
  address: Array<Address>;
}

export interface Channel {
  value: string;
  channelType: string;
  messageType: string;
}
