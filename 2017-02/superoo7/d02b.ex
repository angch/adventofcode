defmodule D02b do
  @moduledoc """
  Documentation for D02b.
  """

  def index do
    fileRead("input.txt")
    |> processFile
    |> main
    |> Enum.reduce(fn a,b -> a + b end)
  end

  @doc """

  ## Examples
      iex> b = "1\\t2\\t3\\n4\\t5\\t6"
      iex> D02b.processFile(b)
      [[1, 2, 3], [4, 5, 6]]
  """
  def main(input) do
    res = input
          |> Enum.map(fn x -> solve(x) end)
      res
  end


  def fileRead(filename) do
    {:ok, input} = File.read(filename)
    input
  end

  def processFile(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn x -> String.split(x,"\t") end)
    |> Enum.map(fn x -> Enum.map(x, fn y -> String.to_integer(y) end) end)
  end

  @doc """

  ## Examples
      iex> D02b.solve([5,9,2,8])
      4
      iex> D02b.solve([9,4,7,3])
      3
      iex> D02b.solve([3,8,6,5])
      2
  """
  def solve(input) do
    for x <- input,
        y <- input,
        x != y,
        rem(x, y) == 0 do
      div(x, y)
    end
    |> Enum.at(0)
  end
end

D02b.index
|> IO.inspect
