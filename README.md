# GoHuffman

Projeto em Go para demonstrar, de forma simples, como funciona o **Código de Huffman**: contagem de frequências, montagem da árvore, geração de códigos binários, codificação, decodificação e cálculo de economia de bits.

## O que é o Código de Huffman

O Código de Huffman é uma técnica de compressão sem perdas que:

- conta a frequência de cada caractere no texto;
- monta uma árvore binária priorizando menores frequências;
- atribui códigos binários menores para caracteres mais frequentes;
- gera uma representação comprimida do texto.

Neste projeto, a ideia é didática: mostrar cada etapa no terminal.

## Estrutura do projeto

- `main.go`: fluxo principal da demonstração e interação via terminal.
- `frequency/counter.go`: contagem de frequência dos caracteres.
- `tree/heap.go` e `tree/build.go`: fila de prioridade (heap) e construção da árvore de Huffman.
- `encoder/encoder.go`: geração dos códigos e codificação do texto.
- `decoder/decoder.go`: reconstrução do texto original a partir do binário.
- `stats/stats.go`: cálculo de bits originais/comprimidos e percentual de economia.
- `utils/printer.go`: impressão da árvore no terminal.
- `model/node.go`: estrutura do nó da árvore.

## Pré-requisitos

- Go `1.22` ou superior.

## Como executar

No diretório do projeto:

```bash
go run .
```

Depois, digite um texto quando solicitado.

## Exemplo rápido

Entrada:

```text
banana
```

Saída (resumo):

- Frequências: `b=1`, `a=3`, `n=2`
- Códigos possíveis: `a=0`, `b=10`, `n=11`
- Texto codificado: `100110110`
- Decodificado: `banana`
- Economia: `81.25%`

## Observações

- A árvore e os códigos podem variar quando há empate de frequências, mas a codificação/decodificação continua correta.
- Este projeto tem foco educacional; não implementa serialização da árvore nem compactação em arquivo binário.
