import gerador from 'k6/x/cpf';
import { sleep } from 'k6'

export const options = {
  vus: 1,
  duration: '3s',

}

export default function () {
  console.log(`Gerando um novo CPF: ${gerador.cpf(false)}`);
  sleep(0.5);
}