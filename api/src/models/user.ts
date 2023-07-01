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

export interface Address {
  zipcode: string;
  street1: string;
  street2: string;
  country: string;
  state: string;
}
