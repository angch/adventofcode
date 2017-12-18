defmodule D01a do
  @moduledoc """
  Documentation for D01a of advent of code.
  """
  def index do
    readFile("input.txt")
    |> main
  end

  @doc """
  Test cases

  ## Examples
      iex> D01a.main("1122")
      3
      iex> D01a.main("1111")
      4
      iex> D01a.main("1234")
      0
      iex> D01a.main("91212129")
      9
  """
  def main(input) do
    output = input
             |> String.split("", trim: true)
             |> Enum.map(&String.to_integer/1)
             |> sum
    output
  end

  def readFile(filename) do
    {:ok, input} = File.read(filename)
    input
  end

  # When 0 -> NONE
  def sum([]), do: 0

  # When there is a list, add h to the end to resolve forth case (AXXXXXXXXA = A)
  def sum([h | _] = list), do: calc_sum(list ++ [h], 0)


  @doc """
  Test cases

  ## Examples
  iex> D01a.calc_sum([], 3)
  3
  iex> D01a.calc_sum([1,2,3], 0)
  0
  iex> D01a.calc_sum([2,3,1,3,1,2,2], 0)
  2
  """

  # When is done, no more list left
  def calc_sum([], total), do: total

  # check if first equal second, if true then add
  def calc_sum([a, a | t], total), do: calc_sum([a | t], total + a)

  # If not match in previous case, then just remove the first
  def calc_sum([_| t], total), do: calc_sum(t, total)
end

D01a.index
|> IO.inspect
