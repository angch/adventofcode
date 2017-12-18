defmodule D02a do
  @moduledoc """
  Documentation for D02a.
  """
  def index do
    fileRead("input.txt")
    |> processFile
    |> main
  end


  @doc """

  ## Examples
      iex> b = "1\\t2\\t3\\n4\\t5\\t6"
      iex> D02a.processFile(b)
      [[1, 2, 3], [4, 5, 6]]
  """
  def processFile(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn x -> String.split(x,"\t") end)
    |> Enum.map(fn x -> Enum.map(x, fn y -> String.to_integer(y) end) end)
  end


  @doc """

  ## Examples
      iex> a = [[1,2,3], [4,5,6]]
      iex> D02a.main(a)
      4
  """
  def main(input) do
    res = input
          |> Enum.map(fn x -> findDifference(x) end)
          |> List.flatten
          |> Enum.reduce(fn (x, y) -> x + y end)
  end

  def fileRead(filename) do
    {:ok, input} = File.read(filename)
    input
  end

  def findDifference(input) do
    [Enum.max(input) - Enum.min(input)]
  end
end

D02a.index
|> IO.inspect
