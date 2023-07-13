import { Address } from './address';

export interface Partner {
  id: string;
  email: string;
  name: string;
  phone: string;
  ext_id: string;
  addresses: Array<Address>;
}
